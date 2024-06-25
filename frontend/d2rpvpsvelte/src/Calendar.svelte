<script>
  import { nickname, csrfToken, getCsrfToken, jwtoken, mode } from "./store.js";
  import { get } from "svelte/store";
  import { onMount } from "svelte";

  export let year;
  export let month;
  export let date;
  export let events;
  export let eventname;
  export let getCalendartext;

  onMount(async () => {
    await getCsrfToken();
  });

  let editing = null;
  let eventText = "";
  let inputRef = null;

  $: ({ daysInMonth, startDayOfWeek } = generateCalendar(year, month));

  $: $nickname; // store의 상태를 구독

  function generateCalendar(year, month) {
    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);

    const daysInMonth = [];
    let currentDay = new Date(firstDay);

    while (currentDay <= lastDay) {
      daysInMonth.push(new Date(currentDay));
      currentDay.setDate(currentDay.getDate() + 1);
    }

    const startDayOfWeek = firstDay.getDay(); // Sunday = 0, Monday = 1, ..., Saturday = 6

    return { daysInMonth, startDayOfWeek };
  }

  $: if (editing && inputRef) {
    inputRef.focus();
  }

  const handleEdit = (day) => {
    if ($nickname === "admin" || $nickname === "admin_m") {
      editing = day.toDateString();
    }
  };

  const handleSave = async (day) => {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const eventDate = day.toDateString();
    const eventPayload = {
      year: day.getFullYear(),
      month: day.getMonth() + 1, // getMonth()는 0부터 시작하므로 1을 더해줍니다
      day: day.getDate(),
      event: eventText,
    };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const endpoint = $mode ? "/changetext_m" : "/changetext";
        const response = await fetch(endpoint, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(eventPayload),
        });

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }

        const data = await response.json();

        // 성공적으로 서버에 데이터가 전송되면, 로컬 상태도 업데이트합니다
        alert("이벤트 기록을 저장하였습니다");

        // 상태를 직접 업데이트
        events = { ...events, [eventDate]: eventText };
        editing = null;

        getCalendartext(year, month);
      } catch (error) {
        console.error("Error:", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  };

  // Create calendar rows
  $: calendarRows = createCalendarRows(daysInMonth, startDayOfWeek);

  function createCalendarRows(daysInMonth, startDayOfWeek) {
    const rows = [];
    let week = [];
    let dayOfWeekCounter = 0;

    // Fill the first week with empty cells if necessary
    for (let i = 0; i < startDayOfWeek; i++) {
      week.push("empty");
      dayOfWeekCounter++;
    }

    // Fill the calendar with the actual days
    daysInMonth.forEach((day) => {
      if (dayOfWeekCounter === 7) {
        rows.push(week);
        week = [];
        dayOfWeekCounter = 0;
      }

      const isToday =
        day.getDate() === date &&
        day.getMonth() === month &&
        day.getFullYear() === year;
      const dayKey = day.getDate();
      week.push({ day, isToday, dayKey });
      dayOfWeekCounter++;
    });

    // Fill the last week with empty cells if necessary
    if (week.length > 0) {
      while (week.length < 7) {
        week.push("empty");
      }
      rows.push(week);
    }

    return rows;
  }

  function handle_eventtext(event) {
    eventText = event.target.value;
  }
</script>

<div class="table-outline">
  <div class="table-head">
    <div class="table-date">일</div>
    <div class="table-date">월</div>
    <div class="table-date">화</div>
    <div class="table-date">수</div>
    <div class="table-date">목</div>
    <div class="table-date">금</div>
    <div class="table-date">토</div>
  </div>
  {#each calendarRows as week}
    <div class="calendar-row">
      {#each week as dayEntry}
        {#if dayEntry === "empty"}
          <div class="table-date empty"></div>
        {:else}
          <div
            class="table-date table-date-heigh"
            style={dayEntry.isToday ? "color: blue;" : ""}
            role="button"
            tabindex="0"
            on:click={() => handleEdit(dayEntry.day)}
            on:keydown={(e) =>
              (e.key === "Enter" || e.key === " ") && handleEdit(dayEntry.day)}
          >
            {dayEntry.day.getDate()}
            <div class="event-cell">
              {#if editing === dayEntry.day.toDateString()}
                <div class="edit-container">
                  <input
                    type="text"
                    bind:this={inputRef}
                    class="calendar-input"
                    bind:value={eventText}
                    on:input={handle_eventtext}
                  /><br />
                  <button on:click={() => handleSave(dayEntry.day)}>수정</button
                  >
                </div>
              {:else}
                <div>
                  {eventname[dayEntry.dayKey] || ""}
                </div>
              {/if}
            </div>
          </div>
        {/if}
      {/each}
    </div>
  {/each}
</div>

<style>
  .calendar-row {
    display: flex;
    width: 100%; /* Ensure rows take full width */
  }

  .table-date {
    width: 14.28%;
    padding: 10px;
    border: 1px solid white;
    align-items: center;
    overflow: hidden; /* Hide overflow text */
    white-space: nowrap; /* Prevent line breaks */
    text-overflow: ellipsis; /* Show ellipsis (...) for overflow text */
  }
  .table-date-heigh {
    height: 100px;
    cursor: pointer;
  }

  .table-date.empty {
    border: 1px solid white;
  }

  .edit-container {
    align-items: center;
  }

  .calendar-input {
    flex-grow: 1;
  }

  .calendar-row:nth-child(even) .table-date:nth-child(even) {
    background-color: #555d;
  }
  .calendar-row:nth-child(odd) .table-date:nth-child(odd) {
    background-color: #555d;
  }
  .calendar-row:nth-child(odd) .table-date:nth-child(even) {
    background-color: #888d;
  }
  .calendar-row:nth-child(even) .table-date:nth-child(odd) {
    background-color: #888d;
  }

  .calendar-input {
    background-color: #0000;
    font-size: 20px;
    color: white;
    border: none;
    font-family: "diablo";
  }
</style>

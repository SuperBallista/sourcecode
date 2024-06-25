<script>
  import { onMount } from "svelte";
  import {
    nicknames,
    fetchNicknames,
    csrfToken,
    getCsrfToken,
    jwtoken,
    mode,
  } from "./store.js";
  import { get } from "svelte/store";

  onMount(async () => {
    await getCsrfToken();
    const currentUrl = window.location.href; // 현재 URL을 가져옵니다
    if (currentUrl.includes("/mpk")) {
      mode.set(true); // URL에 /mpk가 포함되어 있으면 mode를 true로 설정합니다
    } else {
      mode.set(false);
    }

    fetchNicknames($mode); // 닉네임을 가져옴
  });

  let Championship1;
  let Championship2;
  let Championship3;
  let Championship4;

  let Runner_up1;
  let Runner_up2;
  let Runner_up3;
  let Runner_up4;

  let Place3rd1;
  let Place3rd2;
  let Place3rd3;
  let Place3rd4;

  let teamSize;
  let numberteams;
  let eventname;

  function checkForDuplicates() {
    if (teamSize < 2) {
      Championship2 = null;
      Runner_up2 = null;
      Place3rd2 = null;
    }
    if (teamSize < 3) {
      Championship3 = null;
      Runner_up3 = null;
      Place3rd3 = null;
    }
    if (teamSize < 4) {
      Championship4 = null;
      Runner_up4 = null;
      Place3rd4 = null;
    }
    if (numberteams < 4) {
      Runner_up1 = null;
      Runner_up2 = null;
      Runner_up3 = null;
      Runner_up4 = null;
    }
    if (numberteams < 8) {
      Place3rd1 = null;
      Place3rd2 = null;
      Place3rd3 = null;
      Place3rd4 = null;
    }

    const values = [
      Championship1,
      Championship2,
      Championship3,
      Championship4,
      Runner_up1,
      Runner_up2,
      Runner_up3,
      Runner_up4,
      Place3rd1,
      Place3rd2,
      Place3rd3,
      Place3rd4,
    ];

    const filteredValues = values.filter(
      (value) => value != null && value != undefined && value != ""
    );
    const uniqueValues = new Set(filteredValues);
    console.log(filteredValues.length, uniqueValues.size);

    return uniqueValues.size == filteredValues.length;
  }

  function HandleChampionship1(event) {
    Championship1 = event.target.value;
  }

  function HandleChampionship2(event) {
    Championship2 = event.target.value;
  }

  function HandleChampionship3(event) {
    Championship3 = event.target.value;
  }
  function HandleChampionship4(event) {
    Championship4 = event.target.value;
  }

  function HandleRunner_up1(event) {
    Runner_up1 = event.target.value;
  }

  function HandleRunner_up2(event) {
    Runner_up2 = event.target.value;
  }

  function HandleRunner_up3(event) {
    Runner_up3 = event.target.value;
  }
  function HandleRunner_up4(event) {
    Runner_up4 = event.target.value;
  }

  function HandlePlace3rd1(event) {
    Place3rd1 = event.target.value;
  }

  function HandlePlace3rd2(event) {
    Place3rd2 = event.target.value;
  }

  function HandlePlace3rd3(event) {
    Place3rd3 = event.target.value;
  }
  function HandlePlace3rd4(event) {
    Place3rd4 = event.target.value;
  }

  function Handleeventname(event) {
    eventname = event.target.value;
  }

  function Handlenumberteams(event) {
    numberteams = event.target.value;
  }
  function HandleteamSize(event) {
    teamSize = event.target.value;
  }

  async function sendEvent() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const data = {
      Championship1: Championship1,
      Championship2: Championship2,
      Championship3: Championship3,
      Championship4: Championship4,

      Runner_up1: Runner_up1,
      Runner_up2: Runner_up2,
      Runner_up3: Runner_up3,
      Runner_up4: Runner_up4,

      Place3rd1: Place3rd1,
      Place3rd2: Place3rd2,
      Place3rd3: Place3rd3,
      Place3rd4: Place3rd4,

      eventname: eventname,
      numberteams: numberteams,
      teamSize: teamSize,
    };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        if (checkForDuplicates()) {
          const endpoint = $mode ? "/submitevent_m" : "/submitevent";
          const response = await fetch(endpoint, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "CSRF-Token": token,
              Authorization: `Bearer ${checkjwt}`,
            },
            body: JSON.stringify(data), // Send the JSON data as the body of the request
          });
          if (response.ok) {
            alert("대회 기록 전송 완료");
          } else {
            alert("에러 발생");
          }
        } else {
          alert("중복 기록된 선수가 있습니다");
        }
      } catch (error) {
        alert("오류 발생 :", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }
</script>

<div class="left">
  <input
    type="text"
    class="namewidth"
    bind:value={eventname}
    on:change={Handleeventname}
    placeholder="대회 이름"
  />
  <select
    bind:value={numberteams}
    on:change={Handlenumberteams}
    class="namewidth"
  >
    <option value="2">2팀(듀얼)</option>
    <option value="4">3-4팀(4강)</option>
    <option value="8">5-8팀(8강)</option>
    <option value="16">9팀 이상(16강)</option>
    <option value="24">정식전</option>
  </select>
  <select bind:value={teamSize} on:change={HandleteamSize} class="namewidth">
    <option value="1">1인 1팀</option>
    <option value="2">2인 1팀</option>
    <option value="3">3인 1팀</option>
    <option value="4">4인 1팀</option>
  </select>
  우승자
  <select
    bind:value={Championship1}
    on:change={HandleChampionship1}
    class="namewidth"
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>
  <select
    bind:value={Championship2}
    on:change={HandleChampionship2}
    class="namewidth"
    disabled={teamSize < 2}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Championship3}
    on:change={HandleChampionship3}
    class="namewidth"
    disabled={teamSize < 3}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Championship4}
    on:change={HandleChampionship4}
    class="namewidth"
    disabled={teamSize < 4}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  준우승
  <select
    bind:value={Runner_up1}
    on:change={HandleRunner_up1}
    class="namewidth"
    disabled={numberteams < 4}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>
  <select
    bind:value={Runner_up2}
    on:change={HandleRunner_up2}
    class="namewidth"
    disabled={numberteams < 4 || teamSize < 2}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Runner_up3}
    on:change={HandleRunner_up3}
    class="namewidth"
    disabled={numberteams < 4 || teamSize < 3}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Runner_up4}
    on:change={HandleRunner_up4}
    class="namewidth"
    disabled={numberteams < 4 || teamSize < 4}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  3위
  <select
    bind:value={Place3rd1}
    on:change={HandlePlace3rd1}
    class="namewidth"
    disabled={numberteams < 8}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>
  <select
    bind:value={Place3rd2}
    on:change={HandlePlace3rd2}
    class="namewidth"
    disabled={numberteams < 8 || teamSize < 2}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Place3rd3}
    on:change={HandlePlace3rd3}
    class="namewidth"
    disabled={numberteams < 8 || teamSize < 3}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <select
    bind:value={Place3rd4}
    on:change={HandlePlace3rd4}
    class="namewidth"
    disabled={numberteams < 8 || teamSize < 4}
  >
    <option value={null}>선택하세요</option>
    {#each $nicknames as option}
      <option value={option}>{$mode ? option.replace("_m", "") : option}</option
      >
    {/each}
  </select>

  <button on:click={() => sendEvent()}>대회 기록 전송</button>
</div>

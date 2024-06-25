<script>
  let recordData = [];
  let loading = true;
  import { get } from "svelte/store";
  import { onMount } from "svelte";
  import { jwtoken, getCsrfToken, csrfToken } from "./store";

  async function fetchGameData() {
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        // fetch GET 요청을 보냅니다.
        const response = await fetch("/no_approved_record", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
          },
        });

        // 응답을 JSON 형태로 파싱합니다.
        recordData = await response.json();
        loading = false;
      } catch (error) {
        alert("정보를 불러오는 중 오류가 발생하였습니다");
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }

  onMount(async () => {
    fetchGameData();

    await getCsrfToken();
  });

  async function approveRecord(orderNum) {
    // 승인 처리 로직을 추가하세요.
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴
    const data = { orderNum: orderNum };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const response = await fetch("/approve-record", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(data), // Send the JSON data as the body of the request
        });
        if (response.ok) {
          alert("기록을 승인하였습니다");
          fetchGameData();
        }
      } catch (error) {
        alert("에러 발생", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }

  async function deleteRecord(orderNum) {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴
    const data = { orderNum: orderNum };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const response = await fetch("/delete-record", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(data), // Send the JSON data as the body of the request
        });
        if (response.ok) {
          alert("기록을 삭제하였습니다");
          fetchGameData();
        }
      } catch (error) {
        alert("에러 발생", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }
</script>

<table>
  <tr>
    <th class="record-table-date"> 날짜 </th>
    <th class="record-table-loser"> 패자 </th>
    <th class="record-table-score"> 점수 </th>
    <th class="record-table-ok"> 승인 </th>
  </tr>
  {#if loading}<tr> <td colspan="4" class="text-center">로딩 중..</td></tr>{/if}
  {#if !loading && recordData.length == 0}<tr>
      <td colspan="4" class="text-center">기록이 없습니다</td></tr
    >{/if}

  {#each recordData as row}
    <tr>
      <td
        >{new Date(row.Date).toLocaleString("en-US", {
          month: "2-digit",
          day: "2-digit",
        })}</td
      >
      <!-- 적절한 날짜 필드를 JSON에서 가져오세요. -->
      <td>{row.Loser}</td>
      <td>{row.LScore}</td>
      <td>
        <button on:click={() => approveRecord(row.OrderNum)}>승인</button>
        <button on:click={() => deleteRecord(row.OrderNum)}>삭제</button>
      </td>
    </tr>
  {/each}
</table>

<style>
  table {
    width: 90%;
    margin-top: 20px;
  }
  .record-table-date {
    width: 15%;
  }
  .record-table-loser {
    width: 40%;
  }
  .record-table-score {
    width: 15%;
  }
  .record-table-ok {
    width: 20%;
  }
</style>

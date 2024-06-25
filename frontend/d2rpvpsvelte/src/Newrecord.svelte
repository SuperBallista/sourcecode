<script>
  import { onMount } from "svelte";
  import { get } from "svelte/store";

  import {
    csrfToken,
    getCsrfToken,
    jwtoken,
    nicknames,
    fetchNicknames,
    nickname,
    mode,
  } from "./store.js";

  let nicknameexceptme = [];
  let myScore = 0;
  let winner;
  let winnerScore = 5;

  onMount(async () => {
    await getCsrfToken();
    fetchNicknames($mode); // 닉네임을 가져옴
    nicknameexceptme = $nicknames.filter((item) => item !== $nickname);
  });

  async function submitrecord() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const data = {
      winner: winner,
      winnerScore: winnerScore,
      myScore: myScore,
    };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        if (myScore > 4 || myScore < 0) {
          alert("점수 입력이 잘못되었습니다");
          return;
        } else {
          const response = await fetch("/submitrecord", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${checkjwt}`,
              "CSRF-Token": token,
            },
            body: JSON.stringify(data), // Send the JSON data as the body of the request
          });
          if (response.ok) {
            alert("기록 전송 완료");
          } else {
            alert("에러 발생");
          }
        }
      } catch (error) {
        alert("오류 발생 :", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }

  function handle_winner_change(event) {
    winner = event.target.value;
  }
</script>

<div class="left">
  승자 고르기
  <select
    bind:value={winner}
    on:change={handle_winner_change}
    class="namewidth"
  >
    {#each nicknameexceptme as option}
      <option value={option}>{option}</option>
    {/each}
  </select>
  내 점수 입력
  <input type="number" class="scorewidth" bind:value={myScore} />
  <button on:click={() => submitrecord()}>보내기</button>
</div>

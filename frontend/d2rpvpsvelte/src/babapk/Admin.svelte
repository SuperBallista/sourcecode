<script>
  import { onMount } from "svelte";
  import { get } from "svelte/store";

  import {
    nickname,
    csrfToken,
    getCsrfToken,
    jwtoken,
    nicknames,
    fetchNicknames,
    mode,
  } from "../store.js";
  let player;
  let playerscore;
  function handle_player_change(event) {
    player = event.target.value;
  }
  function handle_playerscore_change(event) {
    playerscore = event.target.value;
  }

  onMount(async () => {
    await getCsrfToken();
    fetchNicknames($mode); // 닉네임을 가져옴
  });

  async function submit_bonus_score() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const checkjwt = get(jwtoken);
    if (checkjwt) {
      const data = { player: player, playerScore: playerscore };

      try {
        const response = await fetch("/submit-admin-score", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(data), // Send the JSON data as the body of the request
        });
        if (response.ok) {
          alert("점수 부여 완료");
        } else {
          alert("에러 발생");
        }
      } catch (error) {
        alert("오류 발생 :", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }

  async function score_reset() {
    let userResponse = confirm(
      "모든 참가자 점수를 초기화합니다. 계속하시겠습니까?"
    );

    if (userResponse) {
      const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴
      const checkjwt = get(jwtoken);
      if (checkjwt) {
        try {
          const response = await fetch("/reset-rank", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${checkjwt}`,
              "CSRF-Token": token,
            },
          });
          if (response.ok) {
            alert("점수를 초기화하였습니다");
          } else {
            alert("에러 발생");
          }
        } catch (error) {
          alert("오류 발생 :", error);
        }
      } else {
        alert("다시 로그인해주세요");
      }
    }
  }
</script>

<div class="main_data">
  {#if $nickname === "admin"}
    <div class="left">
      직접 점수를 부여할 선수
      <select
        bind:value={player}
        on:change={handle_player_change}
        class="namewidth"
      >
        {#each $nicknames as option}
          <option value={option}>{option}</option>
        {/each}
      </select>
      <input
        type="number"
        bind:value={playerscore}
        on:change={handle_playerscore_change}
        placeholder="점수"
        class="scorewidth"
      />
      <button on:click={() => submit_bonus_score()}>점수 부여하기</button>
      <br />
      <button on:click={() => score_reset()}>모든 점수 초기화하기</button>
    </div>
  {:else}
    <div class="left">권한이 없습니다.</div>
  {/if}
</div>

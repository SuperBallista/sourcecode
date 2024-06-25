<script>
  import { onMount } from "svelte";
  import { key } from "../store.js";
  let rankData = [];
  let filteredData = [];
  let loading = true;
  let error = null;
  let showDetails = [];
  let selectfilter = "1";

  async function fetchData() {
    try {
      const response = await fetch("/rankdata");
      if (!response.ok) {
        throw new Error("연결 에러입니다");
      }
      rankData = await response.json();
      showDetails = new Array(rankData.length).fill(false); // showDetails 배열 초기화
      filterData(); // 데이터를 가져온 후 필터링 적용
    } catch (err) {
      error = err;
    } finally {
      loading = false;
    }
  }

  function filterData() {
    if (selectfilter === 0) {
      filteredData = rankData;
    } else if (selectfilter === 1) {
      filteredData = rankData.filter((row) => row[6] + row[7] >= 1);
    } else {
      filteredData = rankData.filter((row) => row[6] + row[7] >= selectfilter);
    }
    showDetails = new Array(filteredData.length).fill(false); // 필터링된 데이터에 맞게 showDetails 배열 업데이트
  }

  // 컴포넌트가 마운트될 때 변수를 구독
  onMount(() => {
    const unsubscribe = key.subscribe((value) => {
      fetchData();
    });

    // 컴포넌트가 언마운트될 때 구독 해제
    return () => {
      unsubscribe();
    };
  });

  function toggleDetails(index) {
    showDetails[index] = !showDetails[index];
    showDetails = [...showDetails]; // 배열을 업데이트하여 Svelte가 반응하도록 함
  }

  function handleKeyDown(event, index) {
    if (event.key === "Enter" || event.key === " ") {
      toggleDetails(index);
    }
  }
</script>

<div class="main_data">
  <select
    class="namewidth center"
    bind:value={selectfilter}
    on:change={() => filterData()}
  >
    <option value="0">모든 선수 보기</option>
    <option value="1">전적이 있는 선수만 보기</option>
    <option value="10">전적 10 이상인 선수만 보기</option>
  </select>
</div>
<div class="table-outline main_data rank">
  <div class="table-head">
    <div class="table-head_rank table-contents_cell">등급</div>
    <div class="table-head_nickname table-contents_cell">닉네임</div>
    <div class="table-head_winrate table-contents_cell">승률</div>
    <div class="table-head_score table-contents_right-cell">점수</div>
  </div>
  {#if loading}
    <p>로딩 중...</p>
  {:else if error}
    <p>Error: {error.message}</p>
  {:else}
    {#each filteredData as row, index}
      <div class="table-contents-wrapper">
        <div
          class="table-contents"
          on:click={() => toggleDetails(index)}
          on:keydown={handleKeyDown}
        >
          <div class="table-contents_rank table-contents_cell">
            {#if row[0] / rankData.length <= 0.1}<img
                class="ranksize"
                src="/img/diamond.png"
                alt="10%"
              />
            {:else if row[0] / rankData.length <= 0.25}<img
                class="ranksize"
                src="/img/platinum.png"
                alt="25%"
              />
            {:else if row[0] / rankData.length <= 0.45}<img
                class="ranksize"
                src="/img/gold.png"
                alt="45%"
              />
            {:else if row[0] / rankData.length <= 0.7}<img
                class="ranksize"
                src="/img/silver.png"
                alt="70%"
              />
            {:else}<img class="ranksize" src="/img/bronze.png" alt="100%" />
            {/if}
          </div>
          <!-- 등급 (순위) -->
          <div class="table-contents_nickname table-contents_cell">
            {row[1]}
          </div>
          <!-- 닉네임 -->
          <div class="table-contents_winrate table-contents_cell">
            {row[6] + row[7] == 0
              ? 0
              : Math.round((row[6] / (row[6] + row[7])) * 10000) / 100}
          </div>
          <!-- 승률 (이긴 횟수) -->
          <div class="table-contents_score table-contents_right-cell">
            {row[3]}
          </div>
          <!-- 점수 (BScore + LScore) -->
        </div>
        {#if showDetails[index]}
          <div class="table-contents_detail show">
            대전점수: {row[4]} <br />
            대회점수: {row[5]} <br />
            전적: {row[6] + row[7]} <br />
            승수: {row[6]} <br />
            패수: {row[7]} <br />
            마지막게임: {row[8]}
          </div>
        {/if}
      </div>
    {/each}
  {/if}
</div>

<style>
  .table-head_rank,
  .table-contents_rank {
    width: 15%;
  }

  .table-head_nickname,
  .table-contents_nickname {
    width: 40%;
  }

  .table-head_winrate,
  .table-contents_winrate {
    width: 25%;
  }

  .table-head_score,
  .table-contents_score {
    width: 20%;
  }
</style>

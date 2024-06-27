<script>
  import { onMount } from "svelte";
  import {
    nicknames,
    fetchNicknames,
    nickname,
    key,
    clickformopen,
    csrfToken,
    getCsrfToken,
    jwtoken,
    mode,
  } from "../store.js";
  import { get } from "svelte/store";

  let recordData = [];
  let loading = true;
  let error = null;
  let showDetails = [];
  let currentPage = 1;
  const itemsPerPage = 20;
  let filter = "All";
  let filteredData = [];
  let paginatedData = [];

  async function fetchData() {
    try {
      fetchNicknames($mode); // 닉네임을 가져옴
      const response = await fetch("/recorddata");
      if (!response.ok) {
        throw new Error("연결 에러입니다");
      }
      recordData = await response.json();
      showDetails = new Array(recordData.length).fill(false); // showDetails 배열 초기화
      filteredData = applyFilter(recordData); // 데이터가 로드되면 초기 필터링
      updatePaginatedData(); // 페이지네이션 데이터 초기화
    } catch (err) {
      error = err;
    } finally {
      loading = false;
    }
  }

  // 컴포넌트가 마운트될 때 변수를 구독
  onMount(async () => {
    await getCsrfToken();
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

  function changePage(page) {
    currentPage = page;
    updatePaginatedData(); // 페이지 변경 시 페이지네이션 데이터 업데이트
  }

  function applyFilter(data) {
    if (filter === "All") {
      return data;
    }
    return data.filter(
      (item) => item.Winner === filter || item.Loser === filter
    );
  }

  function updatePaginatedData() {
    const start = (currentPage - 1) * itemsPerPage;
    const end = start + itemsPerPage;
    paginatedData = filteredData.slice(start, end);
  }

  $: totalPages = () => Math.ceil(filteredData.length / itemsPerPage);

  function handleFilterChange(event) {
    filter = event.target.value;
    currentPage = 1; // 필터 변경 시 페이지를 첫 페이지로 초기화
    filteredData = applyFilter(recordData);
    updatePaginatedData(); // 페이지네이션 데이터 초기화
  }

  function visiblePages() {
    const maxVisible = 5; // 한 번에 표시할 최대 페이지 수
    const pages = [];

    if (totalPages() <= maxVisible) {
      for (let i = 1; i <= totalPages(); i++) {
        pages.push(i);
      }
    } else {
      let start = Math.max(currentPage - Math.floor(maxVisible / 2), 1);
      let end = start + maxVisible - 1;

      if (end > totalPages()) {
        end = totalPages();
        start = end - maxVisible + 1;
      }

      for (let i = start; i <= end; i++) {
        pages.push(i);
      }
    }

    return pages;
  }

  function updatekey() {
    key.update((n) => n + 1);
  }

  async function delete_row(OrderNum) {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const data = { OrderNum: OrderNum };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const response = await fetch("/delete-row", {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(data), // Send the JSON data as the body of the request
        });
        if (response.ok) {
          alert("삭제 완료하였습니다");
        } else {
          alert("에러 발생");
        }
      } catch (error) {
        alert("오류 발생 :", error);
      } finally {
        updatekey();
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }
</script>

<div class="filter">
  <select
    class="namewidth center"
    bind:value={filter}
    on:change={handleFilterChange}
  >
    <option value="All">모두 보기</option>
    {#each $nicknames as option}
      <option value={option}>{option}</option>
    {/each}
  </select>
</div>

<div class="table-outline main_data">
  <div class="table-head">
    <div class="table-head_no table-contents_cell">번호</div>
    <div class="table-head_winner table-contents_cell">승자</div>
    <div class="table-head_loser table-contents_cell">패자</div>
    <div class="table-head_date table-contents_right-cell">날짜</div>
  </div>
  {#if loading}
    <p>로딩 중...</p>
  {:else if error}
    <p>Error: {error.message}</p>
  {:else}
    {#each paginatedData as { OrderNum, Winner, Loser, Date, wscore, lscore }, index (index)}
      <div class="table-contents-wrapper">
        <div
          class="table-contents"
          on:click={() =>
            toggleDetails(index + (currentPage - 1) * itemsPerPage)}
          on:keydown={(event) =>
            handleKeyDown(event, index + (currentPage - 1) * itemsPerPage)}
        >
          <div class="table-contents_no table-contents_cell">{OrderNum}</div>
          <div class="table-contents_winner table-contents_cell">{Winner}</div>
          <div class="table-contents_loser table-contents_cell">{Loser}</div>
          <div class="table-contents_date table-contents_right-cell">
            {Date}
          </div>
        </div>
        {#if showDetails[index + (currentPage - 1) * itemsPerPage]}
          <div
            class="table-contents_detail {showDetails[
              index + (currentPage - 1) * itemsPerPage
            ]
              ? 'show'
              : ''}"
          >
            {#if $nickname === "admin" || $nickname === "admin_m"}<button
                on:click={() => delete_row(OrderNum)}>삭제</button
              ><br />{/if}
            승자 점수: {wscore} <br />
            패자 점수: {lscore}
          </div>
        {/if}
      </div>
    {/each}
    <div class="pagination">
      {#if currentPage > 1}
        <button on:click={() => changePage(1)}>First</button>
        <button on:click={() => changePage(currentPage - 1)}>Previous</button>
      {/if}

      {#each visiblePages() as page}
        <button
          class:active={currentPage === page}
          on:click={() => changePage(page)}
        >
          {page}
        </button>
      {/each}

      {#if currentPage < totalPages()}
        <button on:click={() => changePage(currentPage + 1)}>Next</button>
        <button on:click={() => changePage(totalPages())}>Last</button>
      {/if}
    </div>
  {/if}
</div>

{#if $nickname}
  <div class="fixed-button-div">
    <button on:click={() => clickformopen("newrecord")}>등록<br />하기</button
    ><br /><br />
    <button on:click={() => clickformopen("recordok")}>승인<br />하기</button>
  </div>
{/if}

<style>
  .filter {
    margin-bottom: 20px;
    text-align: center;
  }

  .table-head_no,
  .table-contents_no {
    width: 10%;
  }

  .table-head_winner,
  .table-contents_winner {
    width: 35%;
  }

  .table-head_loser,
  .table-contents_loser {
    width: 35%;
  }

  .table-head_date,
  .table-contents_date {
    width: 20%;
  }

  .table-contents {
    border-top: 3px solid white;
    cursor: pointer; /* 클릭 가능한 커서 */
  }

  .table-contents_detail {
    display: none;
    text-align: left;
    padding: 10px;
    border-top: 3px solid white;
    background-color: #222d; /* 세부 정보 배경색 추가 */
  }

  .table-contents_detail.show {
    display: block;
  }

  .pagination {
    display: flex;
    justify-content: center;
    background-color: black;
  }

  .pagination button {
    margin: 0 5px;
    padding: 5px 10px;
    cursor: pointer;
    background-color: #555d;
  }

  .pagination button.active {
    font-weight: bold;
    background-color: #888d;
  }
</style>

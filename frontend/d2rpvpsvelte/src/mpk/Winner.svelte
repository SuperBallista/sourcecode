<script>
  let winners = [
    {
      season: "2024.6 랭킹전",
      first: "스티븐",
      second: "BallistaBB",
      third: "kissa",
      details: "-",
    },
    {
      season: "2024.5 랭킹전",
      first: "BallistaBB",
      second: "스티븐",
      third: "MAX",
      details: "-",
    },
    {
      season: "2024.2 랭킹전",
      first: "Jay",
      second: "제트",
      third: "MAX",
      details: "-",
    },
    {
      season: "2024.1 랭킹전",
      first: "Jay",
      second: "스티븐",
      third: "제트",
      details:
        "23.12 정기전 : 1위 Robe 2위 스티븐 3위 제트<br>23.11 정기전 : 1위 파괴 2위 스티븐 3위 Robe",
    },
    {
      season: "2023.10 랭킹전",
      first: "스티븐",
      second: "달빛",
      third: "스티븐",
      details:
        "정기전 : 1위 CASINO 2위 Queen 3위 드루<br/>23.7 정기전 : 1위 Edward 2위 스티븐 3위 모팔모",
    },
    {
      season: "2023.6 랭킹전",
      first: "스티븐",
      second: "BallistaBB",
      third: "yakuza",
      details: "정기전 : 1위 yakuza 2위 BallistaBB 3위 모팔모",
    },
    {
      season: "2023.5 랭킹전",
      first: "Robe",
      second: "달건",
      third: "Boss",
      details:
        "정기전 : 1위 달건 2위 Robe 3위 Joy<br/>테마전 : 1위 Robe 2위 Southbank 3위 Art_ist",
    },
    {
      season: "2023.10 랭킹전",
      first: "스티븐",
      second: "달빛",
      third: "스티븐",
      details: "정기전 : 1위 CASINO 2위 Queen 3위 드루",
    },
  ];

  let showDetails = Array(winners.length).fill(false);

  function toggleDetails(index) {
    showDetails = showDetails.map((val, i) => (i === index ? !val : val));
  }
</script>

<div class="table-outline main_data">
  <div class="table-head">
    <div class="table-eventname">시즌 이름</div>
    <div class="table-1st">1위</div>
    <div class="table-2nd">2위</div>
    <div class="table-3rd">3위</div>
  </div>
  {#each winners as { season, first, second, third, details }, index}
    <div class="table-contents-wrapper">
      <div class="table-contents" on:click={() => toggleDetails(index)}>
        <div class="table-eventname">{season}</div>
        <div class="table-1st">{first}</div>
        <div class="table-2nd">{second}</div>
        <div class="table-3rd">{third}</div>
      </div>
      {#if showDetails[index]}
        <div class="table-contents_detail show">{@html details}</div>{/if}
    </div>
  {/each}
</div>

<style>
  .table-eventname {
    width: 40%;
    height: 45px;
    text-align: center;
  }
  .table-1st {
    width: 20%;
    height: 45px;
    text-align: center;
  }
  .table-2nd {
    width: 20%;
    height: 45px;
    text-align: center;
  }
  .table-3rd {
    width: 20%;
    height: 45px;
    text-align: center;
  }
  .show {
    text-align: center;
  }
</style>

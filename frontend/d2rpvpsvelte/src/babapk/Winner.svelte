<script>
  let winners = [
    {
      season: "2024.6 랭킹전",
      first: "Head",
      second: "Jay",
      third: "Duky",
      details:
        "3티어 정식전 : 1위 Duky 2위 Jay 3위 IRONKIM<br/>고뇌 정규전 : 1위 Head 2위 jay 3위 봐바<br/>레어 정규전 : 1위 Jay 2위 Duky",
    },
    {
      season: "2024.5 랭킹전",
      first: "Duky",
      second: "Head",
      third: "IRONKIM",
      details:
        "정식전 : 1위 줍줍 2위 Duky 3위 쭈니바바리안<br/>테마전 : 1위 금복이 2위 goyard 3위 IRONKIM",
    },
    {
      season: "2024.4 랭킹전",
      first: "Jay",
      second: "Duky",
      third: "줍줍",
      details:
        "정식전 : 1위 aerts 2위 Head 3위 jay<br/>테마전 : 아리앗옥례 & 오찬물 & Jay",
    },
    {
      season: "2024.3 랭킹전",
      first: "Duky",
      second: "BigMara",
      third: "팍씨",
      details:
        "정식전 : 1위 Joral_Sinsa 2위 Jak 3위 바바<br/>테마전 : 1위 Jay 2위 아이언킴 3위 오찬물 <br/> 테마전 : 1위 Jay 2위 아이언킴 3위 오찬물",
    },
    {
      season: "2024.2 랭킹전",
      first: "barbar",
      second: "오찬물",
      third: "Duky",
      details:
        "정식전 : 1위 오찬물 2위 BallistaBB 3위 GeonDal<br/>테마전 : 1위 옥례 2위 옥땅 3위 IRONKIM",
    },
    {
      season: "2024.1 랭킹전",
      first: "오찬물",
      second: "Duky",
      third: "Head",
      details: "정식전 : 1위 Duky 2위 Jay 3위 오찬물",
    },
  ];

  let Pastwinners = [
    {
      season: "2023.12 랭킹전",
      first: "Duky",
      second: "Jak",
      third: "아리",
      details: "정식전 : 1위 아리 2위 도라라 3위 Duky ",
    },
    {
      season: "2023.11 랭킹전",
      first: "Jak",
      second: "Duky",
      third: "오찬물",
      details:
        "정식전 : 1위 봐바 2위 도라라 3위 Art_ist <br/> 테마전 : 1위 잼재미 2위 Jak 3위 아리앗옥례",
    },
    {
      season: "2023.10 랭킹전",
      first: "오찬물",
      second: "ALO",
      third: "오찬물",
      details:
        "정식전 : 1위 봐바 2위 배로나 3위 잼재미<br/> 테마전 : 아리 & 오찬물 & 아리앗옥례",
    },
    {
      season: "2023.9 랭킹전",
      first: "오찬물",
      second: "Jak",
      third: "바바",
      details: "정식전 : 1위 오찬물 2위 봐바 3위 IRONKIM",
    },
    {
      season: "2023.8 랭킹전",
      first: "Duky",
      second: "Duky",
      third: "IRONKIM",
      details: "정식전 : 1위 Jak 2위 BB 3위 도적",
    },
    {
      season: "2023.7 랭킹전",
      first: "BB(통합우승)",
      second: "BB(고뇌우승)",
      third: "BB(비고뇌우승)",
      details: "정식전 : 1위 Duky 2위 묵선 3위 오찬물",
    },

    {
      season: "시즌 10",
      first: "ALO",
      second: "ADOOBA",
      third: "Duky",
      details: "정식전 10회 : 1위 도적 2위 ALO 3위 Jak",
    },
    {
      season: "시즌 9",
      first: "스티븐",
      second: "ADOOBA",
      third: "Duky",
      details: "정식전 9회 : 1위 Duky 2위 ADOOBA 3위 Jak",
    },
    {
      season: "시즌 8",
      first: "스티븐",
      second: "Duky",
      third: "ALO",
      details:
        "정식전 8회 : 1위 배로나 2위 Duky 3위 Art_ist<br/>테마전 8회 : 1위 GeonDal 2위 BallistaBB 3위 오찬물",
    },
    {
      season: "시즌 7",
      first: "봐바",
      second: "ALO",
      third: "도라라",
      details:
        "정식전 7회 : 1위 PABLO 2위 야토신 3위 DUTY<br/>테마전 7회 : 1위 ADOOBA & 스티븐 & 도라라 2위  3위 ",
    },
    {
      season: "시즌 6",
      first: "PABLO",
      second: "오찬물",
      third: "ADOOBA",
      details:
        "정식전 6회 : 1위 Joral_Sinsa 2위 GeonDal 3위 RFHH<br/>테마전 6회 : 1위 도적 & Jak 2위  3위 ",
    },
    {
      season: "시즌 5",
      first: "ADOOBA",
      second: "오찬물",
      third: "Joral_Sinsa",
      details:
        "정식전 5회 : 1위 ALO 2위 RFHH 3위 TOONGTANG<br/>테마전 5회 : 1위 돌아돌아 & 파멸 2위  3위 ",
    },
    {
      season: "시즌 4",
      first: "Joral_Sinsa",
      second: "오찬물",
      third: "ADOOBA",
      details:
        "정식전 4회 : 1위 파멸 2위 GG 3위 좀건들지마봐<br/>테마전 4회 : 1위 Az 2위 TOONGTANG 3위 행복한우빈이네",
    },
    {
      season: "시즌 3",
      first: "TOONGTANG",
      second: "Joral_Sinsa",
      third: "ALO",
      details:
        "정식전 3회 : 1위 Thanos 2위 MikeTyson 3위 Az<br/>테마전 3회 : 1위 GG 2위 돌아돌아 3위 TOONGTANG",
    },
    {
      season: "시즌 2",
      first: "TOONGTANG",
      second: "BABBARIAN",
      third: "Thanos",
      details:
        "정식전 2회 : 1위 TOONGTANG 2위 AlwaysReview 3위 Az<br/>테마전 2회 : 1위 BABBARIAN & Thanos 2위  3위 ",
    },
    {
      season: "시즌 1",
      first: "TOONGTANG",
      second: "Thanos",
      third: "AlwaysReview",
      details:
        "정식전 10회 : 1위 도적 2위 ALO 3위 Jak<br/>정식전 1회 : 1위 돌아돌아 2위 FBI멀더 3위 BallistaBB<br/>테마전 1회 : 1위 TOONGTANG 2위 BallistaBB 3위 PK",
    },
  ];

  let showDetails = Array(winners.length).fill(false);

  function toggleDetails(index) {
    showDetails = showDetails.map((val, i) => (i === index ? !val : val));
  }

  let showPastDetails = Array(Pastwinners.length).fill(false);
  function togglePastDetails(index) {
    showPastDetails = showPastDetails.map((val, i) =>
      i === index ? !val : val
    );
  }
  let showPastWinner = false;
  function togglePastwinners() {
    showPastWinner = !showPastWinner;
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

  <div class="table-contents-wrapper">
    <div class="table-contents textcenter" on:click={() => togglePastwinners()}>
      <div class="text-center">이전 시즌 결과 보기</div>
    </div>
  </div>
  {#if showPastWinner}
    {#each Pastwinners as { season, first, second, third, details }, index}
      <div class="table-contents-wrapper">
        <div class="table-contents" on:click={() => togglePastDetails(index)}>
          <div class="table-eventname">{season}</div>
          <div class="table-1st">{first}</div>
          <div class="table-2nd">{second}</div>
          <div class="table-3rd">{third}</div>
        </div>
        {#if showPastDetails[index]}
          <div class="table-contents_detail show">{@html details}</div>{/if}
      </div>
    {/each}
  {/if}
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

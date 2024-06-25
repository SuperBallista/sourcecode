<script>
  import { nickname, mode, jwtoken, clickformopen } from "./store";
  import { onMount } from "svelte";
  import { get } from "svelte/store";

  let email = "Loading";
  let tscore = "Loading";
  let bscore = "Loading";
  let lscore = "Loading";
  let date = "Loading";
  let wins = "Loading";
  let loses = "Loading";

  // 사용자 정보가 저장된 세션 데이터를 가져오는 함수 (예시)

  // GET 요청을 보내는 fetch 함수
  async function fetchGameData() {
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const endpoint = $mode ? "/user_data_m" : "/user_data";
        // fetch GET 요청을 보냅니다.

        const response = await fetch(endpoint, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
          },
        });

        // 응답을 JSON 형태로 파싱합니다.
        const data = await response.json();
        email = data.email;
        tscore = data.tscore;
        lscore = data.lscore;
        bscore = data.bscore;
        date = data.lastdate;
        wins = data.countwin;
        loses = data.countlose;
      } catch (error) {
        alert("정보를 불러오는 중 오류가 발생하였습니다");
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }
  onMount(() => {
    fetchGameData();
  });
</script>

<div class="outline">
  <div class="info-flexbox">
    <div class="info-name">닉네임</div>
    <div class="info-data">
      {$mode ? $nickname.replace("_m", "") : $nickname}
    </div>
  </div>
  <div class="info-flexbox">
    <div class="info-name-data">{email}</div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">
      <button on:click={() => clickformopen("changepw")}>암호 변경</button>
    </div>
    <div class="info-data">
      <button on:click={() => clickformopen("changeemail")}>이메일 변경</button>
    </div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">점수총합</div>
    <div class="info-data">{tscore}</div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">대전점수</div>
    <div class="info-data">{bscore}</div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">대회점수</div>
    <div class="info-data">{lscore}</div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">최근경기</div>
    <div class="info-data">{date}</div>
  </div>
  <div class="info-flexbox">
    <div class="info-name">승패</div>
    <div class="info-data">{wins}/{loses}</div>
  </div>
</div>

<style>
  .outline {
    margin-top: 20px;
  }
  .info-flexbox {
    display: flex;
  }
  .info-name {
    width: 300px;
    text-align: center;
    font-size: 20px;
  }
  .info-data {
    width: 300px;
    text-align: center;
    font-size: 20px;
  }
  .info-name-data {
    text-align: center;
    width: 600px;
    font-size: 20px;
  }
</style>

package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CalculationRequest struct {
	MyHP          float64 `json:"myhp"`
	YourHP        float64 `json:"yourhp"`
	MyFrame       int     `json:"myframe"`
	YourFrame     int     `json:"yourframe"`
	MyClass       int     `json:"myclass"`
	YourClass     int     `json:"yourclass"`
	MyCharLv      int     `json:"mycharlv"`
	YourCharLv    int     `json:"yourcharlv"`
	Iterations    int     `json:"iterations"`
	MyAr          float64 `json:"myar"`
	YourAr        float64 `json:"yourar"`
	MyDf          float64 `json:"mydf"`
	YourDf        float64 `json:"yourdf"`
	MyDfoff       float64 `json:"mydfoff"`
	YourDfoff     float64 `json:"yourdfoff"`
	MyDodge       float64 `json:"mydodge"`
	YourDodge     float64 `json:"yourdodge"`
	MyCrush       float64 `json:"mycrush"`
	YourCrush     float64 `json:"yourcrush"`
	MyOpenWound   float64 `json:"myopenwound"`
	YourOpenWound float64 `json:"youropenwound"`
	MyMaxDmg      float64 `json:"mymaxdmg"`
	MyMinDmg      float64 `json:"mymindmg"`
	YourMaxDmg    float64 `json:"yourmaxdmg"`
	YourMinDmg    float64 `json:"yourmindmg"`
	MyReduce      float64 `json:"myreduce"`
	YourReduce    float64 `json:"yourreduce"`
	MyCs          float64 `json:"mycs"`
	MyDs          float64 `json:"myds"`
	YourCs        float64 `json:"yourcs"`
	YourDs        float64 `json:"yourds"`
	MyThorns      float64 `json:"mythorns"`
	YourThorns    float64 `json:"yourthorns"`
}

type CalculationResponse struct {
	WinCount  int     `json:"winCount"`
	LoseCount int     `json:"loseCount"`
	DrawCount int     `json:"drawCount"`
	WinRate   float64 `json:"winRate"`
}

func roundFloatToInt(value float64) float64 {
	return float64(int(math.Floor(value)))
}

func roundFloatToTwoDecimals(value float64) float64 {
	return math.Floor(value*100) / 100
}

func calculateRound(params CalculationRequest) CalculationResponse {
	rand.Seed(time.Now().UnixNano())
	wincount := 0
	losecount := 0

	for number := 0; number < params.Iterations; number++ {
		framecount := 0
		myframecheck := 0
		yourframecheck := 0
		mygamehp := params.MyHP
		yourgamehp := params.YourHP

		mylv := float64(params.MyCharLv)
		if params.MyClass < 3 {
			mylv = float64(params.MyCharLv + 20)
		} else {
			mylv = float64(params.MyCharLv + 5)
		}

		yourlv := float64(params.YourCharLv)
		if params.YourClass < 3 {
			yourlv = float64(params.YourCharLv + 20)
		} else {
			yourlv = float64(params.YourCharLv + 5)
		}

		for mygamehp > 0 && yourgamehp > 0 {
			framecount++            //프레임 진행
			myframecheck++          //내 프레임 진행
			yourframecheck++        // 상대 프레임 진행
			mywoundcondition := 0   // 내 상악 상태 확인
			yourwoundcondition := 0 // 상대 상악 상태 확인

			if myframecheck-params.MyFrame >= 0 { // 내 공격 시점 프레임이 되었을 때
				myframecheck = myframecheck - params.MyFrame
				// 내 프레임 진행 초기화
				r_number := rand.Float64()
				//공격 성공여부 체크랜덤함수 추출
				if r_number < 0.25*roundFloatToTwoDecimals((100-float64(params.YourDodge))/100*params.MyAr/(params.MyAr+params.YourDf*(100-(params.MyDfoff/2))/100)*2*mylv/(mylv+yourlv)) {
					// 내 공격이 성공했는지 여부를 확인
					r_number = float64(rand.Intn(100) + 1) // 내 강타 유효여부 체크랜덤함수 추출
					if r_number < params.MyCrush {         // 내 강타가 유효할 경우
						yourgamehp -= roundFloatToInt(yourgamehp / 10 / params.YourReduce * 100) // 강타 피 적용
					}
					r_number = float64(rand.Intn(100) + 1) // 내 상악 유효여부 체크랜덤함수 추출
					if r_number < params.MyOpenWound {     // 내 상악이 유효할 경우
						yourwoundcondition = 200 // 내 상악 적용
					}
					r_number = roundFloatToInt(float64(rand.Intn(int((params.MyMaxDmg/6*params.YourReduce/100)-(params.MyMinDmg/6*params.YourReduce/100)+1))) + (params.MyMinDmg / 6 * params.YourReduce / 100))
					// 나의 데미지 계산(최소값과 최대값 사이에서)
					cdsrandom := float64(rand.Intn(100) + 1)
					// 내 치타 극대화의 유효여부 체크함수 추출
					if cdsrandom < (params.MyCs + params.MyDs*(100-params.MyCs)/100) {
						// 치타나 극대화가 성공할 경우
						yourgamehp -= r_number // 데미지 한번 더 적용
					}
					yourgamehp -= r_number // 데미지 적용
				}
				mygamehp -= (params.YourThorns / 6 * params.MyReduce / 100) // 나의 쏜즈 데미지 적용
			}

			if yourframecheck-params.YourFrame >= 0 {
				yourframecheck = yourframecheck - params.YourFrame
				r_number := rand.Float64()
				if r_number < 0.25*roundFloatToTwoDecimals((100-float64(params.MyDodge))/100*params.YourAr/(params.YourAr+params.MyDf*(100-(params.YourDfoff/2))/100)*2*yourlv/(yourlv+mylv)) {
					r_number = float64(rand.Intn(100) + 1)
					if r_number < params.YourCrush {
						mygamehp -= roundFloatToInt(mygamehp / 10 / params.MyReduce * 100)
					}
					r_number = float64(rand.Intn(100) + 1)
					if r_number < params.YourOpenWound {
						mywoundcondition = 200
					}
					r_number = roundFloatToInt(float64(rand.Intn(int((params.YourMaxDmg/6*params.MyReduce/100)-(params.YourMinDmg/6*params.MyReduce/100)+1))) + (params.YourMinDmg / 6 * params.MyReduce / 100))
					cdsrandom := float64(rand.Intn(100) + 1)
					if cdsrandom < (params.YourCs + params.YourDs*(100-params.YourCs)/100) {
						mygamehp -= r_number
					}
					mygamehp -= r_number
				}
				yourgamehp -= (params.MyThorns / 6 * params.YourReduce / 100)
			}

			if mywoundcondition > 0 { // 만약 내가 상악 상태라면
				mywoundcondition -= 1                                    // 내 상악상태 시간을 1프레임 줄여줌
				if mygamehp > (45*float64(params.YourCharLv)-1319)/256 { // 내 피가 상악 1프레임 데미지보다 많다면
					mygamehp -= (45*float64(params.YourCharLv) - 1319) / 256 // 상악 피 적용해서 깎임
				} else { // 아니면
					mygamehp = 1 // 피 1로 만듬
				}
			}
			if yourwoundcondition > 0 {
				yourwoundcondition -= 1
				if yourgamehp > (45*float64(params.MyCharLv)-1319)/256 {
					yourgamehp -= (45*float64(params.MyCharLv) - 1319) / 256
				} else {
					yourgamehp = 1
				}
			}
			if mygamehp <= 0 && yourgamehp > 0 { // 내 피가 없고 상대 피는 남은 경우
				losecount += 1 // 패배 1 추가
			}
			if mygamehp > 0 && yourgamehp <= 0 { // 상대 피가 없고 내 피가 남는 경우
				wincount += 1 // 승리 1 추가
			}

			// 무한 루프 방지용 임시 코드 (루프 15000번 돌리면 강제 종료)
			if framecount > 15000 {
				log.Println("Loop count exceeded 15000, breaking to prevent infinite loop")
				break
			}
		}

	}

	drawcount := params.Iterations - wincount - losecount
	WinRate := float64(wincount) / float64(params.Iterations) * 100

	return CalculationResponse{
		WinCount:  wincount,
		LoseCount: losecount,
		DrawCount: drawcount,
		WinRate:   WinRate,
	}
}

func main() {
	router := gin.Default()

	// CORS 설정
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://d2rpvp.org"} // React 앱이 실행되는 도메인
	config.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(config))

	router.POST("/calculate", func(c *gin.Context) {
		startTime := time.Now()

		var req CalculationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("Starting calculation...")
		result := calculateRound(req)
		log.Println("Calculation completed.")

		duration := time.Since(startTime)
		log.Printf("Request processed in %s", duration)

		c.JSON(http.StatusOK, result)
	})

	log.Println("Server is running on 8080 port")
	router.Run(":8080")
}

import JudgeSignup from "../../components/backStage/judge/judgeSignup"
import JudgeInfoChange from "../../components/backStage/judge/judgeInfoChange"
import JudgeAd from "../../components/backStage/judge/judgeAd"
function BackStageJudge() {
  return <div className="grid grid-cols-3 gap-[2rem]">
    <JudgeSignup></JudgeSignup>
    <JudgeInfoChange></JudgeInfoChange>
    <JudgeAd></JudgeAd>
  </div>
}

export default BackStageJudge
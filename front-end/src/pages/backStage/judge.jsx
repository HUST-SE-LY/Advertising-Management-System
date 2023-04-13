import JudgeSignup from "../../components/backStage/judge/judgeSignup"
import JudgeInfoChange from "../../components/backStage/judge/judgeInfoChange"
function BackStageJudge() {
  return <div className="grid grid-cols-3 gap-[2rem]">
    <JudgeSignup></JudgeSignup>
    <JudgeInfoChange></JudgeInfoChange>
  </div>
}

export default BackStageJudge
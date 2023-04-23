import { useNavigate } from "react-router-dom";
import Title from "../backStage/title";
import LongButton from "../longButton";

function Other() {

  const navigate = useNavigate()

  return (
    <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem]">
      <Title title="其他"></Title>
      <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3 ">
        <LongButton content="回到首页" handle={() => navigate("/")}></LongButton>
        <LongButton content="进入应用" handle={() => navigate('/app')}></LongButton>
        <LongButton content="注销企业" color="red"></LongButton>
      </div>
    </div>
  );
}

export default Other;

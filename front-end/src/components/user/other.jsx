import { useNavigate } from "react-router-dom";
import Title from "../backStage/title";
import LongButton from "../longButton";
import useAxios from "../../utils/useAxios";

function Other() {
  const navigate = useNavigate();
  const axios = useAxios();
  async function logout() {
    const res = await axios.post("/company/logout");
    localStorage.removeItem("token");
    navigate('/');
  }
  async function cancel() {
    const res = await axios.post("/company/cancel");
  }
  return (
    <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem] animate-listItemIn">
      <Title title="其他"></Title>
      <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3 ">
        <LongButton content="回到首页" handle={() => navigate("/")}></LongButton>
        <LongButton content="进入应用" handle={() => navigate('/app')}></LongButton>
        <LongButton content="退出登录" handle={logout}></LongButton>
        <LongButton content="注销企业" color="red" handle={cancel}></LongButton>
      </div>
    </div>
  );
}

export default Other;

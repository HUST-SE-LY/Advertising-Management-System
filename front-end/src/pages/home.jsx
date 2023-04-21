import { useEffect } from "react";
import NavigateButton from "../components/navigateButton";
import TypingWord from "../components/typingWord";
import useAxios from "../utils/useAxios";
function Home() {
  const axios = useAxios()
  useEffect(() => {
    (async () => {
      const res = await axios.post("/company/register",{})
      console.log(res)
    })()

  },[])


  return (
    <div className="w-screen h-screen grid grid-cols-[3fr_2fr] justify-center items-center bg-slate-50">
      <div className="w-[400px] h-[400px] rounded-full bg-blue-200/50 fixed top-[50px] left-[50px] animate-float"></div>
      <div className="left flex flex-col gap-[3rem] relative">
        <div className="w-[200px] h-[200px] rounded-full bg-blue-300/50 right-[45%] bottom-[50px] fixed animate-[float_8s_infinite]"></div>
        <TypingWord word={"Welcome"}></TypingWord>
        <div className="flex px-[5rem]">
          <NavigateButton
            content={"管理员登录"}
            url={"/manager-login"}
          ></NavigateButton>
          <NavigateButton
            content={"客户登录"}
            url={"/company-login"}
          ></NavigateButton>
          <NavigateButton content={"进入应用"} url={"/app"}></NavigateButton>
        </div>
      </div>
      <div className="right bg-[url('/src/assets/homeBack.jpg')] w-full h-full bg-center bg-cover shadow-blue-300 shadow-home"></div>
    </div>
  );
}

export default Home;

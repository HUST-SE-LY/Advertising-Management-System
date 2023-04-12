import NavigateButton from "../components/navigateButton";
import TypingWord from "../components/typingWord";
function Home() {
  return (
    <div className="w-screen h-screen grid grid-cols-[3fr_2fr] justify-center items-center bg-slate-50">
      <div className="left flex flex-col gap-[3rem]">
        <TypingWord word={"Welcome"}></TypingWord>
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
      <div className="right bg-[url('src/assets/homeBack.jpg')] w-full h-full bg-center bg-cover shadow-blue-300 shadow-home"></div>
    </div>
  );
}

export default Home;

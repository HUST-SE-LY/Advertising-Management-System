import { useEffect } from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import Input from "../components/input";
import LongButton from "../components/longButton";
function CompanyLogin() {
  const [companyId, setCompanyId] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="w-screen h-screen flex justify-center items-center bg-[url('src/assets/companyLoginBack.jpg')] bg-cover bg-center">
      <div className="grid grid-cols-2 w-1/2 h-1/2 bg-white rounded-lg p-[1rem] gap-[2rem] animate-floatin">
        <div className=" bg-slate-200 p-5 flex flex-col box-border rounded">
          <p className="font-sans font-bold text-[2rem] leading-[3rem]">
            Put your AD on our app
          </p>
          <p className="font-sans text-[1.2rem] leading-[2rem]">Sign in or sign up</p>
          <p className="font-sans text-[1.2rem] leading-[2rem]">
            and choose an Ad to put
          </p>
          <div className="w-full h-fit max-h-[10rem] flex-1 mt-[4rem] bg-blue-300 rounded-lg shadow-lg shadow-blue-300/50 p-[1rem] flex flex-col justify-center">
            <p className="leading-[2rem]">
              成为我们的客户之后，你可以在我们的应用上购买广告位，展示您想展示的内容
            </p>
            <p className="text-right">——hust-se-ly</p>
          </div>
        </div>
        <div className="flex flex-col justify-center gap-[1rem] pr-[2rem]">
          <p className="text-[1.5rem] font-bold">企业客户登录</p>
          <Input title="账号" setInfo={setCompanyId}></Input>
          <Input title="密码" setInfo={setPassword}></Input>
          <LongButton content="登录"></LongButton>
          <p className="text-sm text-blue-400 underline cursor-pointer text-right transition-all hover:text-blue-600"><Link to="/company-signup">没有账号？注册</Link></p>
        </div>
      </div>
    </div>
  );
}

export default CompanyLogin;
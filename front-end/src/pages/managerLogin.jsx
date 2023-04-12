import { useEffect } from "react";
import { useState } from "react";
import Input from "../components/input";
import LongButton from "../components/longButton";
function ManagerLogin() {
  const [managerId, setManagerId] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="w-screen h-screen flex justify-center items-center bg-[url('src/assets/managerLoginBack.jpg')] bg-cover bg-center">
      <div className="grid grid-cols-2 w-1/2 h-1/2 bg-white rounded-lg p-[1rem] gap-[2rem] animate-floatin">
        <div className=" bg-slate-200 p-5 flex flex-col box-border rounded">
          <p className="font-sans font-bold text-[2rem] leading-[3rem]">
            Let me help you manager AD.
          </p>
          <p className="font-sans text-[1.2rem] leading-[2rem]">Join us</p>
          <p className="font-sans text-[1.2rem] leading-[2rem]">
            and be the manager
          </p>
          <div className="w-full h-fit flex-1 mt-[4rem] bg-blue-300 rounded-lg shadow-lg shadow-blue-300/50 p-[1rem] flex flex-col justify-center">
            <p className="leading-[2rem]">
              成为管理员后，你可以管理客户，管理广告投放，以及其他有意思的事
            </p>
            <p className="text-right">——hust-se-ly</p>
          </div>
        </div>
        <div className="flex flex-col justify-center gap-[1rem] pr-[2rem]">
          <p className="text-[1.5rem] font-bold">管理员登录</p>
          <Input title="账号" setInfo={setManagerId}></Input>
          <Input title="密码" setInfo={setPassword}></Input>
          <LongButton content="登录"></LongButton>
        </div>
      </div>
    </div>
  );
}

export default ManagerLogin;

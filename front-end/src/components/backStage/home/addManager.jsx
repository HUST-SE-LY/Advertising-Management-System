import Title from "../title";
import Input from "../../input";
import { useState } from "react";
import LongButton from "../../longButton";
import useAxios from "../../../utils/useAxios";
function AddManager(props) {
  const axios = useAxios();
  const [account, setAccount] = useState("");
  const [password, setPassword] = useState("");
  const [secondPassword, setSecondPassword] = useState("");
  async function createManager() {
    if(password !== secondPassword) {
      alert("两次输入的密码不同！");
      return;
    }
    const res = axios.post("/manage/admin/create",{
      account: account,
      password: password,
    })
    props.close();
  }
  return (
    <div
      className="animate-fadein fixed top-0 left-0 w-screen h-screen bg-gray-300/40 z-30 cursor-pointer flex justify-center items-center"
      onClick={props.close}
    >
      <div
        className="w-1/2 flex flex-col gap-[1rem] bg-white cursor-default rounded-xl shadow-xl shadow-gray-200 py-[2rem] px-[5rem]"
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <Title title="添加管理员"></Title>
        <Input title="账号" setInfo={setAccount}></Input>
        <Input title="密码" setInfo={setPassword} ></Input>
        <Input title="再次输入密码" setInfo={setSecondPassword}></Input>
        <LongButton content="添加" handle={createManager}></LongButton>
      </div>
    </div>
  );
}

export default AddManager;

import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Input from "../components/input";
import LongButton from "../components/longButton";
import useAxios from "../utils/useAxios";

function CompanySignup(props) {
  const [account, setAccount] = useState("");
  const [password, setPassword] = useState("");
  const [secondPassword, setSecondPassword] = useState("");
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");
  const [manager, setManager] = useState("");
  const [managerTel, setManagerTel] = useState("");
  const [businessLicenseNumber, setBusinessLicenseNumber] = useState("");
  const [position, setPosition] = useState("-translate-x-0");
  const navigate = useNavigate();
  const axios = useAxios();
  function nextStep() {
    if (password === secondPassword) {
      setPosition("-translate-x-full");
    } else {
      alert("输入密码不一致！");
    }
  }
  function lastStep() {
    setPosition("-translate-x-0");
  }
  async function finalStep() {
    try {
      await axios.post("/company/register", {
        account: account,
        password: password,
        name: name,
        address: address,
        manager_name: manager,
        manager_tel: managerTel,
        business_license_number: businessLicenseNumber,
      });
      setPosition("-translate-x-[200%]");
    } catch (err) {
      console.log(err);
    }
  }
  function goHome() {
    navigate("/");
  }
  return (
    <div className="w-screen h-screen flex justify-center items-center bg-[url('/src/assets/companyLoginBack.jpg')] bg-cover bg-center">
      <div className="grid grid-cols-[1fr_2fr] w-1/2 h-1/2 bg-white rounded-lg p-[1rem] gap-[2rem] animate-floatin">
        <div className="bg-blue-200 w-full h-full rounded-lg p-[1rem] flex flex-col justify-center">
          <p className="text-center text-[1.5rem] font-bold mb-[1rem]">
            注册流程
          </p>
          <div
            className={`transition-all text-center py-[1rem] ${
              position === "-translate-x-0"
                ? "bg-blue-600 text-white rounded-2xl shadow-lg shadow-blue-300"
                : "bg-none"
            }`}
          >
            网站信息
          </div>
          <div className="w-[2px] h-[2rem] bg-blue-400 mx-[auto] my-[0.5rem]"></div>
          <div
            className={`transition-all text-center py-[1rem] ${
              position === "-translate-x-full"
                ? "bg-blue-600 text-white rounded-2xl shadow-lg shadow-blue-300"
                : "bg-none"
            }`}
          >
            企业信息
          </div>
          <div className="w-[2px] h-[2rem] bg-blue-400 mx-[auto] my-[0.5rem]"></div>
          <div
            className={`transition-all text-center py-[1rem] ${
              position === "-translate-x-[200%]"
                ? "bg-blue-600 text-white rounded-2xl shadow-lg shadow-blue-300"
                : "bg-none"
            }`}
          >
            成功申请
          </div>
        </div>
        <div className="flex flex-col justify-center overflow-hidden">
          <div className={`flex ${position} transition-all`}>
            <div className={`w-full flex-shrink-0 flex flex-col gap-[1rem]`}>
              <Input title={"账号"} setInfo={setAccount}></Input>
              <Input
                type={"password"}
                setInfo={setPassword}
                title={"密码"}
              ></Input>
              <Input
                type={"password"}
                setInfo={setSecondPassword}
                title={"再次输入密码"}
              ></Input>
              <LongButton content="下一步" handle={nextStep}></LongButton>
            </div>
            <div className={"w-full flex-shrink-0 flex flex-col gap-[1rem]"}>
              <div className="flex gap-[1rem] w-full">
                <Input title={"企业名称"} setInfo={setName}></Input>
                <Input title={"企业地址"} setInfo={setAddress}></Input>
              </div>
              <div className="flex gap-[1rem] w-full">
                <Input title={"负责人姓名"} setInfo={setManager}></Input>
                <Input title={"负责人电话 +86"} setInfo={setManagerTel}></Input>
              </div>
              <Input
                title={"营业执照号码"}
                setInfo={setBusinessLicenseNumber}
              ></Input>
              <LongButton content="上一步" handle={lastStep}></LongButton>
              <LongButton content="申请" handle={finalStep}></LongButton>
            </div>
            <div className={`w-full flex-shrink-0 flex flex-col gap-[1rem]`}>
              <p className="text-[1.5rem]">申请成功！</p>
              <p>审核后即可正常登录！</p>
              <LongButton content="返回首页" handle={goHome}></LongButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default CompanySignup;

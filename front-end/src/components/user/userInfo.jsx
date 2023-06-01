import { useState } from "react";
import Title from "../backStage/title";
import LongButton from "../longButton";
import NoBorderInput from "../noBorderInput";
import useAxios from "../../utils/useAxios";

function UserInfo(props) {

  const [showEditInfo,setShowEditInfo] = useState(false);
  const [showEditPassword, setShowEditPassword] = useState(false);

  return (
    <>
    <div className=" row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem] animate-listItemIn">
      <Title title="个人信息"></Title>
      <div className="flex flex-col gap-[1rem] mt-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue">
        <p>账号：{props.info.account}</p>
        <p>企业名称：{props.info.name}</p>
        <p>企业地址：{props.info.address}</p>
        <p>负责人：{props.info.manager_name}</p>
        <p>负责人电话：{props.info.manager_tel}</p>
        <p>营业执照：{props.info.business_license_number}</p>
        <LongButton handle={() => setShowEditInfo(true)} content="修改信息"></LongButton>
        <LongButton handle={() => setShowEditPassword(true)} content="修改密码"></LongButton>
      </div>
    </div>
    {
      showEditInfo?<EditInfo info={props.info} hide={() => setShowEditInfo(false)}></EditInfo>:null
    }
    {
      showEditPassword?<EditPassword hide={() => setShowEditPassword(false)}></EditPassword>:null
    }


    </>
    
  );
}


function EditInfo(props) {
  const axios = useAxios();
  const [name, setName] = useState(props.info.name);
  const [address, setAddress] = useState(props.info.address);
  const [managerName, setManagerName] = useState(props.info.manager_name);
  const [managerTel, setManagerTel] = useState(props.info.manager_tel);
  const [businessLicenseNumber, setBusinessLicenseNumber] = useState(props.info.business_license_number);
  async function changeInfo() {
    const res = await axios.post("/company/update-info",{
      account: props.info.account,
      address: address,
      name: name,
      manager_name: managerName,
      manager_tel: managerTel,
      business_license_number:businessLicenseNumber
    })
    console.log(res)
  }
  return <div className="flex items-center justify-center w-full h-full bg-white/20 absolute z-10  top-0 left-0 animate-fadein" onClick={props.hide}>
    <div className="p-[2rem] w-1/2 h-1/2 bg-white rounded-2xl shadow-2xl shadow-gray-500/20 overflow-y-auto scrollbar-blue" onClick={(e) => e.stopPropagation()}>
      <Title title="修改信息"></Title>
      <NoBorderInput value={name} setInfo={setName}>公司名称</NoBorderInput>
      <NoBorderInput value={address} setInfo={setAddress}>公司地址</NoBorderInput>
      <NoBorderInput value={managerName} setInfo={setManagerName}>负责人</NoBorderInput>
      <NoBorderInput value={managerTel} setInfo={setManagerTel}>负责人电话</NoBorderInput>
      <NoBorderInput value={businessLicenseNumber} setInfo={setBusinessLicenseNumber}>营业执照</NoBorderInput>
      <LongButton content="提交申请" handle={changeInfo}></LongButton>
    </div>
  </div>
}

function EditPassword(props) {
  const axios = useAxios()
  const [password, setPassword] = useState("");
  const [secondPassword, setSecondPassword] = useState("");
  async function changePassword() {
    if(password !== secondPassword) {
      alert("两次输入的密码不一致！")
      return ;
    }
    const res = await axios.post("/company/update-pwd",{
      password: password,
    })
    if(res.data.code === 200) {
      alert("修改成功！");
      setPassword("");
      setSecondPassword("")
    }
  }
  return <div className="flex items-center justify-center w-full h-full bg-white/20 absolute z-10  top-0 left-0 animate-fadein" onClick={props.hide}>
    <div className="p-[2rem] w-1/2 h-1/2 bg-white rounded-2xl shadow-2xl shadow-gray-500/20 overflow-y-auto scrollbar-blue" onClick={(e) => e.stopPropagation()}>
      <Title title="修改密码"></Title>
      <NoBorderInput type="password" value={password} setInfo={setPassword}>输入新密码</NoBorderInput>
      <NoBorderInput type="password" value={secondPassword} setInfo={setSecondPassword}>再次输入新密码</NoBorderInput>
      <LongButton content="修改" handle={changePassword}></LongButton>
    </div>
  </div>
}



export default UserInfo;

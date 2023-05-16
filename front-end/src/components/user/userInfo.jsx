import { useState } from "react";
import Title from "../backStage/title";
import LongButton from "../longButton";
import NoBorderInput from "../noBorderInput";

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
        <p>负责人：{props.info.managerName}</p>
        <p>负责人电话：{props.info.managerTel}</p>
        <p>营业执照：{props.info.businessLicenseNumber}</p>
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
  return <div className="flex items-center justify-center w-full h-full bg-white/20 absolute z-10  top-0 left-0 animate-fadein" onClick={props.hide}>
    <div className="p-[2rem] w-1/2 h-1/2 bg-white rounded-2xl shadow-2xl shadow-gray-500/20 overflow-y-auto scrollbar-blue" onClick={(e) => e.stopPropagation()}>
      <Title title="修改信息"></Title>
      <NoBorderInput value={props.info.name}>公司名称</NoBorderInput>
      <NoBorderInput value={props.info.address}>公司地址</NoBorderInput>
      <NoBorderInput value={props.info.managerName}>负责人</NoBorderInput>
      <NoBorderInput value={props.info.managerTel}>负责人电话</NoBorderInput>
      <NoBorderInput value={props.info.businessLicenseNumber}>营业执照</NoBorderInput>
      <LongButton content="提交申请"></LongButton>
    </div>
  </div>
}

function EditPassword(props) {
  return <div className="flex items-center justify-center w-full h-full bg-white/20 absolute z-10  top-0 left-0 animate-fadein" onClick={props.hide}>
    <div className="p-[2rem] w-1/2 h-1/2 bg-white rounded-2xl shadow-2xl shadow-gray-500/20 overflow-y-auto scrollbar-blue" onClick={(e) => e.stopPropagation()}>
      <Title title="修改密码"></Title>
      <NoBorderInput type="password">输入新密码</NoBorderInput>
      <NoBorderInput type="password">再次输入新密码</NoBorderInput>
      <LongButton content="修改"></LongButton>
    </div>
  </div>
}



export default UserInfo;

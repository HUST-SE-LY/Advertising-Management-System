import Title from "../backStage/title";
import LongButton from "../longButton";


function UserInfo(props) {
  return <div className=" row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem]">
    <Title title="个人信息"></Title>
    <div className="flex flex-col gap-[1rem] mt-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue">
      <p>账号：{props.info.account}</p>
      <p>企业名称：{props.info.name}</p>
      <p>企业地址：{props.info.address}</p>
      <p>负责人：{props.info.managerName}</p>
      <p>负责人电话：{props.info.managerTel}</p>
      <p>营业执照：{props.info.businessLicenseNumber}</p>
      <LongButton content="修改信息"></LongButton>
      <LongButton content="修改密码"></LongButton>
    </div>
  </div>
}

export default UserInfo;
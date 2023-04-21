import { useState } from "react"
import UserInfo from "../components/user/userInfo"
import AdInfo from "../components/user/adInfo"


function User() {
  const [userInfo, setUserInfo] = useState({
    account: "1",
    name: "A",
    address: "china",
    managerName: "me",
    managerTel: "123456",
    businessLicenseNumber: "1111111111",
  })




  return <div>
    <TabBar></TabBar>
    <div className="grid grid-rows-2 grid-cols-3 w-[100vw] h-[calc(100vh_-_4rem)] p-[2rem] gap-[2rem]">
      <UserInfo info={userInfo}></UserInfo>
      <AdInfo></AdInfo>
    </div>
  </div>
}

function TabBar() {
  return <div className="w-full px-[2rem] py-[1rem] flex items-center rounded-[0_0_2rem_2rem] gap-[2rem] bg-blue-300 shadow-lg shadow-blue-600/10 ">
    <p className="text-2xl font-bold text-gray-700 tracking-widest">欢迎回来!</p>
  </div>
}


export default User
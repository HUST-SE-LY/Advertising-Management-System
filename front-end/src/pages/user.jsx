import { useState } from "react"
import UserInfo from "../components/user/userInfo"
import AdInfo from "../components/user/adInfo"
import AccountInfo from "../components/user/accountInfo"
import WaitingList from "../components/user/waitingList"
import MyAd from "../components/user/myAd"
import Other from "../components/user/other"


function User() {
  const [userInfo, setUserInfo] = useState({
    account: "1",
    name: "A",
    address: "china",
    managerName: "me",
    managerTel: "123456",
    businessLicenseNumber: "1111111111",
  })

  const [accountInfo, setAccountInfo] = useState({
    money: 114514,
    record: [
      {
        id: 0,
        time: "2023-3-25",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
      {
        id: 1,
        time: "2023-3-23",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
      {
        id: 2,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },{
        id: 3,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },{
        id: 4,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
    ],
  })




  return <div>
    <TabBar></TabBar>
    <div className="relative grid grid-rows-2 grid-cols-3 w-[100vw] h-[calc(100vh_-_4rem)] p-[2rem] gap-[2rem]">
      <UserInfo info={userInfo}></UserInfo>
      <AdInfo></AdInfo>
      <AccountInfo info={accountInfo}></AccountInfo>
      <WaitingList></WaitingList>
      <MyAd></MyAd>
      <Other></Other>
    </div>
  </div>
}

function TabBar() {
  return <div className="w-full px-[2rem] py-[1rem] flex items-center rounded-[0_0_2rem_2rem] gap-[2rem] bg-blue-300 shadow-lg shadow-blue-600/10 ">
    <p className="text-2xl font-bold text-gray-700 tracking-widest">欢迎回来!</p>
  </div>
}



export default User
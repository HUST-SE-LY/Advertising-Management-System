import { useNavigate } from "react-router-dom";

function AccountInfo(props) {
  return <div className="bg-gray-100 rounded-xl p-[2rem] h-full animate-floatin">
    <p className="ml-[1rem] text-2xl font-bold tracking-wide relative before:absolute before:left-[-10px] before:w-[5px] before:h-full before:top-0 before:bg-blue-400">账户信息</p>
    <p className="my-[1rem] text-lg">账户余额：<span className="text-blue-600">{props.balance}</span>元</p>
    <p>持有广告：</p>
    <div className="flex mt-[1rem] flex-col gap-[1rem] overflow-y-auto scrollbar-blue pr-3">
      {
        props.info.Ad.map((info) => <SingleAdInfo key={info.position} info={info}></SingleAdInfo>)
      }
    </div>
  </div>
}

function SingleAdInfo(props) {

  const navigate = useNavigate()
  return <div className="flex gap-[1rem] items-center bg-blue-100 rounded-2xl p-[1rem] animate-listItemIn">
    <p>{props.info.position}号广告位</p>
    <p>剩余{props.info.leftTime}天</p>
    <button className="px-[1rem] py-[0.5rem] rounded-2xl bg-blue-500 text-white hover:shadow-lg hover:shadow-blue-600/20 hover:bg-blue-600 transition-all ml-[auto]" onClick={() => navigate(`/back-stage/ads/${props.info.position}`)}>详情</button>
  </div>
}

export default AccountInfo;
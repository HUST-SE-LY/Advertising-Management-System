import { useState } from "react";
import Title from "../backStage/title";
import LongButton from "../longButton";
import NoBorderInput from "../noBorderInput";

function AccountInfo(props) {

  const [showChargeBoard, setShowChargeBoard] = useState(false)

  return (
    <>
      <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem] animate-listItemIn">
        <Title title="账户信息"></Title>
        <p className="h-[2rem] leading-[2rem]">
          账户余额：<span className="text-blue-500">{props.info.money}</span>元
        </p>
        <div className="h-[calc(100%_-_7.5rem)] flex flex-col gap-[1rem] overflow-auto scrollbar-blue pr-3">
          <p>消费记录</p>
          <div className="flex gap-[1rem] flex-shrink-0">
            <p className="w-[7rem] text-center">时间</p>
            <p className="w-[5rem] text-center">花费/元</p>
            <p className="w-[5rem] text-center">购买广告位</p>
            <p className="w-[4rem] text-center">购买时长</p>
            <p className="w-full flex-1 text-center">订单号</p>
          </div>
          {props.info.record.map((info) => (
            <div className="flex gap-[1rem] bg-blue-100 rounded-xl py-[1rem]">
              <p className="w-[7rem] text-center">{info.time}</p>
              <p className="w-[5rem] text-center">{info.price}</p>
              <p className="w-[5rem] text-center">{info.AdPosition}</p>
              <p className="w-[4rem] text-center">{info.AdTime}</p>
              <p className="w-full flex-1 text-center">{info.id}</p>
            </div>
          ))}
        </div>
        <LongButton content="充值" handle={() => setShowChargeBoard(true)}></LongButton>
      </div>
      {
        showChargeBoard ? <ChargeBoard hide={() => setShowChargeBoard(false)}></ChargeBoard> : null
      }
    </>
  );
}

function ChargeBoard(props) {
  return <div className="flex items-center justify-center w-full h-full bg-white/20 absolute z-10  top-0 left-0 animate-fadein" onClick={props.hide}>
  <div className="p-[2rem] w-1/2 h-1/2 bg-white rounded-2xl shadow-2xl shadow-gray-500/20 overflow-y-auto scrollbar-blue" onClick={(e) => e.stopPropagation()}>
    <Title title="充值"></Title>
    <NoBorderInput>充值金额/元</NoBorderInput>
    <LongButton content="充值"></LongButton>
  </div>
</div>
}

export default AccountInfo;

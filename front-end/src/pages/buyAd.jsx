import { useState } from "react";
import { useParams } from "react-router-dom";
import Title from "../components/backStage/title";

function BuyAd() {
  const params = useParams();
  const [info, setInfo] = useState({
    id: 1,
    state: "not free",
    price: 100,
    expires: "2023-5-1",
  });

  return (
    <div className="w-screen h-screen grid grid-cols-2 gap-[2rem] p-[2rem]">
      <div className="rounded-2xl bg-gray-50 p-[2rem] mt-[1rem]">
        <Title title="广告详情"></Title>
        <div className="flex flex-col gap-[1rem]">
          <p>广告id：{info.id + 1}</p>
          <p>
            广告状态：
            <span
              className={`${
                info.state === "free" ? "text-blue-600" : "text-red-600"
              }`}
            >
              {info.state === "free" ? "空闲" : "占用"}
            </span>
          </p>
        </div>
      </div>
      <div className="rounded-2xl bg-gray-50"></div>
    </div>
  );
}

export default BuyAd;

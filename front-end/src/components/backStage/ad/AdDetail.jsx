import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Title from "../title";
import Input from "../../input";
import LongButton from "../../longButton";

function AdDetail(props) {
  const params = useParams();
  const [info, setInfo] = useState(null);
  async function getInfo() {
    setInfo({
      state: "using",
      price: 100,
      owner: "AAA",
      time: "20",
      url: "https://www.google.com",
      image: "https://w.wallhaven.cc/full/g7/wallhaven-g77oz7.jpg",
    });
  }
  useEffect(() => {
    getInfo();
  }, [params]);

  return (
    <>
      {info ? (
        <>
          <ShowInfo info={info}></ShowInfo>
          <SetInfo info={info}></SetInfo>
        </>
      ) : (
        <></>
      )}
    </>
  );
}

function ShowInfo(props) {
  return (
    <div className="bg-white min-h-[calc(50%_-_3.5rem)]  flex flex-col gap-[1rem] rounded-[2rem] shadow-xl shadow-gray-200/50 p-[2rem]">
      <Title title="广告栏信息"></Title>
      <div className="h-full overflow-y-auto scrollbar-blue animate-fadein">
        <p className=" leading-[3rem] text-lg relative before:w-[3px] before:h-full before:left-[-10px] before:top-0 before:bg-blue-300 before:absolute">
          状态：
          {props.info.state === "using" ? (
            <span className="text-red-400">使用中</span>
          ) : (
            <span className="text-red-400">空闲</span>
          )}
        </p>
        <p className="leading-[3rem] text-lg relative before:w-[3px] before:h-full before:left-[-10px] before:top-0 before:bg-blue-300 before:absolute">
          价格：<span className="text-blue-500">{props.info.price}</span>元/天
        </p>
        {props.info.state === "using" ? (
          <>
            <p className="leading-[3rem] text-lg relative before:w-[3px] before:h-full before:left-[-10px] before:top-0 before:bg-blue-300 before:absolute">
              所属公司：
              <span className="text-blue-500">{props.info.owner}</span>
            </p>
            <p className="leading-[3rem] text-lg relative before:w-[3px] before:h-full before:left-[-10px] before:top-0 before:bg-blue-300 before:absolute">
              剩余时长：
              <span className="text-blue-500">{props.info.price}</span>天
            </p>
            <p className="leading-[3rem] text-lg relative before:w-[3px] before:h-full before:left-[-10px] before:top-0 before:bg-blue-300 before:absolute">
              链接：
              <span className="text-blue-500">
                <a>{props.info.url}</a>
              </span>
            </p>
            <img src={props.info.image} alt="" className="rounded-xl w-5/6" />
          </>
        ) : (
          <></>
        )}
      </div>
    </div>
  );
}

function SetInfo(props) {
  return (
    <div className="animate-fadein bg-white min-h-[calc(50%_-_3.5rem)] rounded-[2rem] shadow-xl shadow-gray-200/50 p-[2rem]">
      <Title title="编辑广告栏"></Title>
      <div className="h-[calc(100%_-_2rem)] flex flex-col mt-[1rem] gap-[1rem] overflow-y-auto scrollbar-blue">
        <Input title="编辑价格 元/天"></Input>

        {
          props.info.state === 'using'?<LongButton color="red" content="停止展示"></LongButton>:<></>
        }
        <LongButton content="确认修改"></LongButton>
      </div>
    </div>
  );
}

export default AdDetail;

import { useRef, useState } from "react";
import { useParams } from "react-router-dom";
import Title from "../components/backStage/title";
import PcModel from "../components/backStage/ad/pcModel";
import PhoneModel from "../components/backStage/ad/phoneModel";
import Input from "../components/input";
import LongButton from "../components/longButton";

function BuyAd() {
  const params = useParams();
  console.log(params);
  const [url, setUrl] = useState("");
  const [title, setTitle] = useState("");
  const [time, setTime] = useState("");
  const fileInput = useRef(null);
  const [picture, setPicture] = useState(null);

  const [info, setInfo] = useState({
    id: parseInt(params.id),
    state: "not free",
    price: 100,
    expires: "2023-5-1",
  });

  async function upload() {
    if (time !== parseInt(time).toString() || parseInt(time) === 0) {
      alert("时间必须填写有效的数字！");
      return ;
    }
    if(url && time && title && picture) {
      const formData = new FormData();
      formData.append("")
    } else {
      alert("请填写完整信息！")
    }

  }

  return (
    <div className="w-screen h-screen grid grid-cols-2 gap-[2rem] p-[2rem]">
      <div className="rounded-2xl h-[calc(100vh_-_4rem)] scrollbar-blue bg-gray-50 overflow-y-auto pr-3 p-[2rem]">
        <Title title="广告详情"></Title>
        <div className="flex flex-col gap-[1rem] my-[1rem]">
          <p>广告id：{info.id + 1}</p>
          <p>
            广告定价：<span className="text-blue-600">{info.price}/天</span>
          </p>
          <p>
            广告状态：
            <span
              className={`${
                info.state === "free" ? "text-blue-600" : "text-red-600"
              }`}
            >
              {info.state === "free" ? "空闲" : "已经被申请"}
            </span>
          </p>
          {info.state === "free" ? null : (
            <p>
              当前广告位已经被申请，您申请的广告在审批通过后将在其他广告（若审批通过）过期后被展示
            </p>
          )}
          {info.state === "free" ? (
            <p>
              申请通过后将<span className="text-blue-500">立即</span>被展示
            </p>
          ) : (
            <p>
              申请通过后最迟开始展示的日期：
              <span className="text-blue-500">{info.expires}</span>
            </p>
          )}
          <p>审批通过后将自动从余额中扣除存款</p>
        </div>
        <Title title="填写信息"></Title>
        <div className="flex flex-col gap-[1rem] my-[1rem]">
          <p>
            <span className="text-red-500 mr-2">*</span>填写要跳转的链接url
          </p>
          <Input info={url} setInfo={setUrl}></Input>
          <p>
            <span className="text-red-500 mr-2">*</span>填写购买时长/天
          </p>
          <Input info={time} setInfo={setTime}></Input>
          <p>
            <span className="text-red-500 mr-2">*</span>填写广告标题
          </p>
          <Input info={title} setInfo={setTitle}></Input>
          <p>
            <span className="text-red-500 mr-2">*</span>选择展示图片
          </p>
          <input
            accept="image/*"
            ref={fileInput}
            className="hidden"
            type="file"
            onChange={() => {
              setPicture(fileInput.current.files[0]);
            }}
          ></input>
          <button
            className="transition-all w-fit h-fit px-[2rem] py-[0.5rem] bg-blue-500 text-white rounded-lg hover:bg-blue-600 hover:shadow-xl hover:shadow-blue-500/10"
            onClick={() => {
              fileInput.current.click();
            }}
          >
            {picture ? "选择完毕" : "点击选择图片"}
          </button>
          <LongButton content="提交申请" handle={upload}></LongButton>
        </div>
      </div>
      <div className="grid grid-rows-2 gap-[1rem] h-[calc(100vh_-_6rem)]">
        <div className="grid grid-cols-2 ">
          <div className="p-[2rem] flex flex-col gap-[1rem] text-lg justify-center">
            <p className="text-right text-2xl text-gray-700 font-bold">
              广告位预览
            </p>
            <p className="text-right">下为PC端预览</p>
            <p className="text-right">右为移动端预览</p>
          </div>
          <PhoneModel></PhoneModel>
        </div>
        <PcModel></PcModel>
      </div>
    </div>
  );
}

export default BuyAd;

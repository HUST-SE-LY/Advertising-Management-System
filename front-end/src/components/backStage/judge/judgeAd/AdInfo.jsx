import LongButton from "../../../longButton";

function AdInfo(props) {
  return (
    <div
      className=" animate-fadein fixed top-0 left-0 w-screen h-screen bg-gray-300/40 z-30 cursor-pointer flex justify-center items-center"
      onClick={props.close}
    >
      <div
        className="w-1/2  bg-white cursor-default rounded-xl shadow-xl shadow-gray-200 py-[2rem] px-[5rem]"
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <p className="mb-[1rem] text-2xl text-gray-700 font-bold tracking-wider">广告信息</p>
        <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 mt-[1rem]">企业名称：<span className="mt-[0.5rem] text-lg text-black">{props.info.company_account}</span></p>
        <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 mt-[1rem]">广告标题：<span className="mt-[0.5rem] text-lg text-black">{props.info.title}</span></p>
        <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 mt-[1rem]">广告图片：</p>
        <img className="rounded-lg mt-[1rem] mb-[1rem] max-w-full max-h-[20rem] shadow-lg" src={`https://aivwe.top${props.info.image_url.split("aivwe.top/.")[1]}`} alt="" />
        <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 text-lg">广告url:<a className="underline text-blue-600" target="_blank" href={props.info.jump_to_url}>{props.info.jump_to_url}</a></p>
        <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 mt-[1rem] text-lg">{props.info.position}号广告位</p>
        <LongButton content="通过" handle={props.pass}></LongButton>
        <LongButton content="拒绝" color="red" handle={props.reject}></LongButton>
      </div>
    </div>
  );
}

export default AdInfo;

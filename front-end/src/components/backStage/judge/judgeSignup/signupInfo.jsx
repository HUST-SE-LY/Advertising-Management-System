import LongButton from "../../../longButton";

function SignupInfo(props) {
  return (
    <div
      className="animate-fadein fixed top-0 left-0 w-screen h-screen bg-gray-300/40 z-30 cursor-pointer flex justify-center items-center"
      onClick={props.close}
    >
      <div
        className="w-1/2 bg-white cursor-default rounded-xl shadow-xl shadow-gray-200 py-[2rem] px-[5rem] "
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <p className="text-center text-2xl text-gray-700 font-bold tracking-wider">
          企业信息
        </p>
        <div className="flex text-lg mb-[1rem]">
          <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
            企业名称：
          </p>
          <span>{props.info.name}</span>
        </div>
        <div className="flex text-lg mb-[1rem]">
          <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
            企业地址：
          </p>
          <span>{props.info.address}</span>
        </div>
        <div className="flex text-lg mb-[1rem]">
          <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
            负责人：
          </p>
          <span>{props.info.manager_name}</span>
        </div>
        <div className="flex text-lg mb-[1rem]">
          <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
            负责人电话：
          </p>
          <span>{props.info.manager_tel}</span>
        </div>
        <div className="flex text-lg mb-[1rem]">
          <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
            营业执照：
          </p>
          <span>{props.info.business_license_number}</span>
        </div>
        <LongButton content="通过" handle={props.pass}></LongButton>
        <LongButton
          content="拒绝"
          color="red"
          handle={props.reject}
        ></LongButton>
      </div>
    </div>
  );
}

export default SignupInfo;

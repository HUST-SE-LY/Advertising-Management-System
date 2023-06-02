import LongButton from "../../../longButton";

function ChangedInfoDetail(props) {
  return (
    <div
      className="animate-fadein fixed top-0 left-0 w-screen h-screen bg-gray-300/40 z-30 cursor-pointer flex justify-center items-center"
      onClick={props.close}
    >
      <div
        className="w-2/3  bg-white cursor-default rounded-xl shadow-xl shadow-gray-200 py-[2rem] px-[5rem] grid grid-cols-2 gap-[2rem] justify-center"
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <div>
          <p className="mb-[2rem] text-2xl text-gray-700 font-bold tracking-wider">
            修改前企业信息
          </p>
          <div className="flex text-lg mb-[1.5rem]">
            <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
              企业名称：
            </p>
            <span>{props.info.previous_name}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
              企业地址：
            </p>
            <span>{props.info.previous_address}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
              负责人：
            </p>
            <span>{props.info.previous_manager_name}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
              负责人电话：
            </p>
            <span>{props.info.previous_manager_tel}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <p className="relative before:absolute before:w-[3px] before:h-full before:top-0 before:left-[-10px] before:bg-blue-300 w-[10rem] ">
              营业执照：
            </p>
            <span>{props.info.previous_business_license_number}</span>
          </div>
        </div>
        <div>
          <p className="mb-[2rem] text-2xl text-gray-700 font-bold tracking-wider">
            修改后企业信息
          </p>
          <div className="flex text-lg mb-[1.5rem]">
            <span className={`w-full px-2 rounded ${props.info.new_name === props.info.previous_name?'':'bg-blue-200'}`}>{props.info.new_name}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <span className={`w-full px-2 rounded ${props.info.new_address === props.info.previous_address?'':'bg-blue-200'}`}>{props.info.new_address}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <span className={`w-full px-2 rounded ${props.info.new_manager_name === props.info.previous_manager_name?'':'bg-blue-200'}`}>{props.info.new_manager_name}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <span className={`w-full px-2 rounded ${props.info.new_manager_tel === props.info.previous_manager_tel?'':'bg-blue-200'}`}>{props.info.new_manager_tel}</span>
          </div>
          <div className="flex text-lg mb-[1.5rem]">
            <span className={`w-full px-2 rounded ${props.info.new_business_license_number === props.info.previous_business_license_number?'':'bg-blue-200'}`}>{props.info.new_business_license_number}</span>
          </div>
        </div>
        <LongButton content="通过" handle={props.pass}></LongButton>
        <LongButton content="拒绝" color="red" handle={props.reject}></LongButton>
      </div>
    </div>
  );
}

export default ChangedInfoDetail;

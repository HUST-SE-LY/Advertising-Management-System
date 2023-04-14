function BasicInfo(props) {
  return (
    <div className="bg-gray-100 rounded-xl p-[2rem] h-full animate-floatin">
      <p className="ml-[1rem] text-2xl font-bold tracking-wide relative before:absolute before:left-[-10px] before:w-[5px] before:h-full before:top-0 before:bg-blue-400">
        基本信息
      </p>
      <div className="h-[calc(100%_-_3rem)] text-lg mt-[1rem] scrollbar-blue overflow-y-auto ">
        <div className="ml-[1rem] relative before:absolute before:left-[-10px] before:w-[2px] before:h-full before:top-0 before:bg-blue-300 before:z-10">
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">公司名称：</p>
            <span>{props.info.name}</span>
          </div>
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">公司账号：</p>
            <span>{props.info.account}</span>
          </div>
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">公司地址：</p>
            <span>{props.info.address}</span>
          </div>
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">公司负责人：</p>
            <span>{props.info.managerName}</span>
          </div>
          
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">负责人电话：</p>
            <span>{props.info.managerTel}</span>
          </div>
          <div className="my-[0.5rem] flex">
            <p className="w-[10rem]">营业执照：</p>
            <span>{props.info.businessLicenseNumber}</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default BasicInfo;

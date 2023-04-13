function Title(props) {
  return <p className="w-fit relative before:absolute before:w-[4px] before:h-full before:z-20 before:bg-blue-600/50 before:left-[-10px] before:top-0 text-2xl font-bold track-wide text-gray-600 hover:before:w-[5px] before:transition-all after:absolute after:w-[0px] after:h-[1px] after:left-0 after:bottom-[-2px] after:transition-all hover:after:w-full after:z-10 after:bg-blue-600/50">{props.title}</p>
}

export default Title
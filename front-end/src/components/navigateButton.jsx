import { Link } from "react-router-dom";

function NavigateButton(props) {
  return (
    <Link to={props.url} className="w-fit mx-[auto]">
      <div className="font-sans w-48 h-16 bg-blue-300/30 text-center text-xl leading-[4rem] rounded-lg hover:bg-blue-600/80 hover:text-white transition-all mx-[auto] animate-floatin">{props.content}</div>
    </Link>
  );
}

export default NavigateButton;

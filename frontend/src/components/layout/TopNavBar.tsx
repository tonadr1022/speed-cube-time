import { RxHamburgerMenu, RxTimer } from "react-icons/rx";
import { FaGear, FaListUl } from "react-icons/fa6";
import { blurElement } from "../../util/handleDropdownClick";
import { Link } from "react-router-dom";
import clsx from "clsx";

type TopNavItemProps = {
  icon: any;
  text: string;
  href: string;
};
const TopNavMenuItem = ({ icon, text, href }: TopNavItemProps) => {
  return (
    <li>
      <Link className="flex" to={href} onClick={blurElement}>
        {icon}
        <span>{text}</span>
      </Link>
    </li>
  );
};

type Props = {
  className?: string;
};
const TopNavBar = ({ className }: Props) => {
  return (
    <div className={clsx("navbar bg-base-200 min-h-8", className)}>
      <div className="navbar-start">
        <div className="dropdown">
          <label tabIndex={0} className="btn btn-sm btn-ghost btn-circle">
            <RxHamburgerMenu />
          </label>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box"
          >
            <TopNavMenuItem icon={<RxTimer />} text="Timer" href="/" />
            {/* <TopNavMenuItem icon={<FaChartPie />} text="Stats" href="/stats" /> */}
            {/* <TopNavMenuItem */}
            {/*   icon={<FaCubesStacked />} */}
            {/*   text="Solves" */}
            {/*   href="solves" */}
            {/* /> */}
            <TopNavMenuItem
              icon={<FaListUl />}
              text="Sessions"
              href="sessions"
            />
            <TopNavMenuItem icon={<FaGear />} text="Settings" href="settings" />
          </ul>
        </div>
      </div>
      <div className="navbar-center">
        <img
          src={"/pwa-512x512.png"}
          alt="ChronoCube Logo"
          width={30}
          height={30}
        />
      </div>
    </div>
  );
};

export default TopNavBar;

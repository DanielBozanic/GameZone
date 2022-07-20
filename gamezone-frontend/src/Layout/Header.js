import logo from "../Assets/images/logo.PNG";
import "../Assets/css/header.css";
import { Link } from "react-router-dom";
import { Navbar, NavbarBrand } from "reactstrap";
import UnauthenticatedNavbar from "./UnauthenticatedNavbar";
import UserNavbar from "./UserNavbar";
import EmployeeNavbar from "./EmployeeNavbar";
import AdminNavbar from "./AdminNavbar";
import * as authService from "../Auth/AuthService";
import * as role from "../Utils/Role";

const Header = () => {
	const navbar = () => {
		if (authService.getRole() === role.ROLE_USER) {
			return <UserNavbar />;
		} else if (authService.getRole() === role.ROLE_EMPLOYEE) {
			return <EmployeeNavbar />;
		} else if (authService.getRole() === role.ROLE_ADMIN) {
			return <AdminNavbar />;
		} else {
			return <UnauthenticatedNavbar />;
		}
	};

	return (
		<Navbar className="header" color="dark" expand="md" dark>
			<NavbarBrand>
				<Link to="/">
					<img src={logo} alt="Logo" className="responsive-img-header" />
				</Link>
			</NavbarBrand>
			{navbar()}
		</Navbar>
	);
};

export default Header;

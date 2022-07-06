import logo from "../Assets/images/logo.PNG";
import "../Assets/css/header.css";
import { Link } from "react-router-dom";
import { Navbar, NavbarBrand } from "reactstrap";

const Header = () => {
	return (
		<Navbar className="header" color="dark" expand="md" dark>
			<NavbarBrand>
				<Link to="/">
					<img src={logo} alt="Logo" className="responsive-img-header" />
				</Link>
			</NavbarBrand>
		</Navbar>
	);
};

export default Header;

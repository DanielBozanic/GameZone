import logo from "../Assets/images/logo.PNG";
import "../Assets/css/footer.css";
import { Navbar, NavbarBrand } from "reactstrap";
import { Link } from "react-router-dom";

const Footer = () => {
	return (
		<Navbar className="footer" color="dark" expand="xl" dark>
			<NavbarBrand>
				<Link to="/">
					<img src={logo} alt="Logo" className="responsive-img-footer" />
				</Link>
			</NavbarBrand>
		</Navbar>
	);
};

export default Footer;

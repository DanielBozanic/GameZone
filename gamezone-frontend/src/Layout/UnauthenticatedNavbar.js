import { Nav, NavbarToggler, Collapse, NavItem, NavLink } from "reactstrap";
import { useState } from "react";
import { Link } from "react-router-dom";

const UnauthenticatedNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Nav className="me-auto" navbar>
					<NavItem>
						<NavLink>
							<Link to="/signUp">Sign Up</Link>
						</NavLink>
					</NavItem>
					<NavItem>
						<NavLink>
							<Link to="/signIn">Sign In</Link>
						</NavLink>
					</NavItem>
				</Nav>
			</Collapse>
		</>
	);
};

export default UnauthenticatedNavbar;

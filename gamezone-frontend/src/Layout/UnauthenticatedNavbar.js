import {
	Nav,
	NavbarToggler,
	Collapse,
	NavItem,
	NavLink,
	Container,
	Col,
} from "reactstrap";
import { useState } from "react";
import { Link } from "react-router-dom";

const UnauthenticatedNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Container fluid style={{ padding: "0" }}>
					<Nav className="me-auto" navbar>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/signUp">Sign Up</Link>
								</NavLink>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/signIn">Sign In</Link>
								</NavLink>
							</NavItem>
						</Col>
						<Col md="auto">
							<NavItem>
								<NavLink>
									<Link to="/viewNews">News</Link>
								</NavLink>
							</NavItem>
						</Col>
					</Nav>
				</Container>
			</Collapse>
		</>
	);
};

export default UnauthenticatedNavbar;

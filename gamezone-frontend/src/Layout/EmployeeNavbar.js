import {
	Nav,
	NavbarToggler,
	Collapse,
	NavItem,
	NavLink,
	UncontrolledDropdown,
	DropdownToggle,
	DropdownMenu,
	DropdownItem,
} from "reactstrap";
import { useState } from "react";
import { Link } from "react-router-dom";
import * as authService from "../Auth/AuthService";
import "../Assets/css/navbar.css";

const EmployeeNavbar = () => {
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);

	const signOut = () => {
		authService.removeToken();
	};

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Nav className="me-auto" navbar>
					<NavItem>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								Add new product
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/addNewVideoGame">
										Video game
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/addNewConsole">
										Console
									</Link>
								</DropdownItem>
								<UncontrolledDropdown inNavbar nav>
									<DropdownToggle caret nav>
										Hardware components
									</DropdownToggle>
									<DropdownMenu right>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewGraphicsCard">
												Graphics card
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewProcessor">
												Processor
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewMotherboard">
												Motherboard
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewRam">
												RAM
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewSsd">
												Solid state drive
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewHdd">
												Hard disk drive
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewMonitor">
												Monitor
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewPsu">
												Power supply unit
											</Link>
										</DropdownItem>
									</DropdownMenu>
								</UncontrolledDropdown>
								<UncontrolledDropdown inNavbar nav>
									<DropdownToggle caret nav>
										Accessories
									</DropdownToggle>
									<DropdownMenu right>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewHeadphones">
												Headphones
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewKeyboard">
												Keyboard
											</Link>
										</DropdownItem>
										<DropdownItem>
											<Link className="drop-down-link" to="/addNewMouse">
												Mouse
											</Link>
										</DropdownItem>
									</DropdownMenu>
								</UncontrolledDropdown>
							</DropdownMenu>
						</UncontrolledDropdown>
					</NavItem>
					<NavItem>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								News
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/addNewsArticle">
										Add news article
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link className="drop-down-link" to="/">
										View news
									</Link>
								</DropdownItem>
							</DropdownMenu>
						</UncontrolledDropdown>
					</NavItem>
					<NavItem>
						<NavLink href="/signIn" onClick={signOut}>
							Sign Out
						</NavLink>
					</NavItem>
				</Nav>
			</Collapse>
		</>
	);
};

export default EmployeeNavbar;

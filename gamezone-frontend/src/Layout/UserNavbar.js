import {
	Nav,
	NavbarToggler,
	UncontrolledDropdown,
	DropdownToggle,
	DropdownMenu,
	DropdownItem,
	Collapse,
	NavItem,
	NavLink,
} from "reactstrap";
import { useState, useEffect } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import { Link } from "react-router-dom";
import * as authService from "../Auth/AuthService";
import * as newsSubscriptionAPI from "../APIs/NewsMicroservice/news_subscription_api";

toast.configure();
const UserNavbar = () => {
	const customId = "UserNavbar";
	const [collapsed, setCollapsed] = useState(true);
	const toggleNavbar = () => setCollapsed(!collapsed);
	const [subscribed, setSubscribed] = useState(false);

	useEffect(() => {
		isUserSubscribed();
	}, []);

	const isUserSubscribed = () => {
		axios
			.get(`${newsSubscriptionAPI.IS_USER_SUBSCRIBED}`)
			.then((res) => {
				setSubscribed(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const subscribe = () => {
		axios
			.post(`${newsSubscriptionAPI.SUBSCRIBE}`)
			.then((res) => {
				setSubscribed(true);
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const unsubscribe = () => {
		axios
			.delete(`${newsSubscriptionAPI.UNSUBSCRIBE}`)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				setSubscribed(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const signOut = () => {
		authService.removeToken();
	};

	return (
		<>
			<NavbarToggler onClick={toggleNavbar} className="mr-2" />
			<Collapse isOpen={!collapsed} navbar>
				<Nav className="me-auto" navbar>
					<NavItem>
						<NavLink>
							<Link to="/profile">Profile</Link>
						</NavLink>
					</NavItem>
					<NavItem>
						<NavLink>
							<Link to="/shoppingCart">Shopping Cart</Link>
						</NavLink>
					</NavItem>
					<NavItem>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								News
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/viewNews">
										News
									</Link>
								</DropdownItem>
								<DropdownItem>
									{!subscribed && (
										<Link className="drop-down-link" to="#" onClick={subscribe}>
											Subscribe
										</Link>
									)}
									{subscribed && (
										<Link
											className="drop-down-link"
											to="#"
											onClick={unsubscribe}
										>
											Unsubscribe
										</Link>
									)}
								</DropdownItem>
							</DropdownMenu>
						</UncontrolledDropdown>
					</NavItem>
					<NavItem>
						<UncontrolledDropdown inNavbar nav>
							<DropdownToggle caret nav>
								Contact Center
							</DropdownToggle>
							<DropdownMenu right>
								<DropdownItem>
									<Link className="drop-down-link" to="/contact">
										Send contact message
									</Link>
								</DropdownItem>
								<DropdownItem>
									<Link
										className="drop-down-link"
										to={`/contactMessages/${authService.getId()}`}
									>
										Contact messages
									</Link>
								</DropdownItem>
							</DropdownMenu>
						</UncontrolledDropdown>
					</NavItem>
					<NavItem>
						<NavLink href="/signIn" onClick={signOut}>
							Sign out
						</NavLink>
					</NavItem>
				</Nav>
			</Collapse>
		</>
	);
};

export default UserNavbar;

import "../../Assets/css/filter.css";
import {
	Button,
	Modal,
	ModalHeader,
	ModalBody,
	ModalFooter,
	Form,
	FormGroup,
	Label,
	Col,
} from "reactstrap";
import { useState, useEffect } from "react";
import Select from "react-select";
import { useForm, Controller } from "react-hook-form";
import axios from "axios";
import * as mouseAPI from "../../APIs/ProductMicroservice/mouse_api";

const MouseFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [dpis, setDpis] = useState([]);
	const [connections, setConnections] = useState([]);

	useEffect(() => {
		getManufacturers();
		getDpis();
		getConnections();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Dpis: [],
			Connections: [],
		},
	});

	const onFilterClick = (filterData) => {
		toggle();
		props.onFilterClick && props.onFilterClick(filterData);
	};

	const filter = (filterData) => {
		onFilterClick(filterData);
	};

	const getManufacturers = () => {
		axios
			.get(mouseAPI.GET_MANUFACTURERS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setManufacturers(temp);
				setValue("Manufacturers", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getDpis = () => {
		axios
			.get(mouseAPI.GET_DPIS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setDpis(temp);
				setValue("Dpis", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getConnections = () => {
		axios
			.get(mouseAPI.GET_CONNECTIONS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setConnections(temp);
				setValue("Connections", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<Col md="2">
				<Button className="filter-button" onClick={toggle}>
					Filter
				</Button>
			</Col>

			<Modal isOpen={modal} toggle={toggle}>
				<ModalHeader toggle={toggle}>Filter</ModalHeader>
				<ModalBody>
					<Form>
						<FormGroup>
							<Label>Filter by manufacturer</Label>
							<Controller
								name={"Manufacturers"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={manufacturers}
											isMulti={true}
											onChange={(val) => {
												onChange(
													manufacturers
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={manufacturers.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by DPI</Label>
							<Controller
								name={"Dpis"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={dpis}
											isMulti={true}
											onChange={(val) => {
												onChange(
													dpis
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={dpis.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by connection</Label>
							<Controller
								name={"Connections"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={connections}
											isMulti={true}
											onChange={(val) => {
												onChange(
													connections
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={connections.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
					</Form>
				</ModalBody>
				<ModalFooter>
					<Button className="confirm-filter-btn" onClick={handleSubmit(filter)}>
						Filter
					</Button>
				</ModalFooter>
			</Modal>
		</>
	);
};

export default MouseFilter;

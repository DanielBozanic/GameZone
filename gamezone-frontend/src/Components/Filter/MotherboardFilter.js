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
import * as motherboardAPI from "../../APIs/ProductMicroservice/motherboard_api";

const MotherboardFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [sockets, setSockets] = useState([]);
	const [processorTypes, setProcessorTypes] = useState([]);
	const [formFactors, setFormFactors] = useState([]);

	useEffect(() => {
		getManufacturers();
		getSockets();
		getProcessorTypes();
		getFormFactors();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Sockets: [],
			ProcessorTypes: [],
			FormFactors: [],
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
			.get(motherboardAPI.GET_MANUFACTURERS)
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

	const getSockets = () => {
		axios
			.get(motherboardAPI.GET_SOCKETS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setSockets(temp);
				setValue("Sockets", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getProcessorTypes = () => {
		axios
			.get(motherboardAPI.GET_PROCESSOR_TYPES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setProcessorTypes(temp);
				setValue("ProcessorTypes", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getFormFactors = () => {
		axios
			.get(motherboardAPI.GET_FORM_FACTORS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setFormFactors(temp);
				setValue("FormFactors", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<Col md="2">
				<Button className="my-button filter-button" onClick={toggle}>
					Filter
				</Button>
			</Col>

			<Modal isOpen={modal} toggle={toggle}>
				<ModalHeader className="title" toggle={toggle}>
					Filter
				</ModalHeader>
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
							<Label>Filter by socket</Label>
							<Controller
								name={"Sockets"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={sockets}
											isMulti={true}
											onChange={(val) => {
												onChange(
													sockets
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={sockets.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by processor type</Label>
							<Controller
								name={"ProcessorTypes"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={processorTypes}
											isMulti={true}
											onChange={(val) => {
												onChange(
													processorTypes
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={processorTypes.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by form factor</Label>
							<Controller
								name={"FormFactors"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={formFactors}
											isMulti={true}
											onChange={(val) => {
												onChange(
													formFactors
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={formFactors.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
					</Form>
				</ModalBody>
				<ModalFooter>
					<Button className="my-button" onClick={handleSubmit(filter)}>
						Filter
					</Button>
				</ModalFooter>
			</Modal>
		</>
	);
};

export default MotherboardFilter;

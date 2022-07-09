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
import * as ramAPI from "../../APIs/ProductMicroservice/ram_api";

const RAMFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [capacities, setCapacities] = useState([]);
	const [memoryTypes, setMemoryTypes] = useState([]);
	const [speeds, setSpeeds] = useState([]);

	useEffect(() => {
		getManufacturers();
		getCapacities();
		getMemoryTypes();
		getSpeeds();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Capacities: [],
			MemoryTypes: [],
			Speeds: [],
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
			.get(ramAPI.GET_MANUFACTURERS)
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

	const getCapacities = () => {
		axios
			.get(ramAPI.GET_CAPACITIES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setCapacities(temp);
				setValue("Capacities", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getMemoryTypes = () => {
		axios
			.get(ramAPI.GET_MEMORY_TYPES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setMemoryTypes(temp);
				setValue("MemoryTypes", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getSpeeds = () => {
		axios
			.get(ramAPI.GET_SPEEDS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setSpeeds(temp);
				setValue("Speeds", res.data);
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
				<ModalHeader className="filter-modal-title" toggle={toggle}>
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
							<Label>Filter by capacity</Label>
							<Controller
								name={"Capacities"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={capacities}
											isMulti={true}
											onChange={(val) => {
												onChange(
													capacities
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={capacities.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by memory type</Label>
							<Controller
								name={"MemoryTypes"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={memoryTypes}
											isMulti={true}
											onChange={(val) => {
												onChange(
													memoryTypes
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={memoryTypes.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by speed</Label>
							<Controller
								name={"Speeds"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={speeds}
											isMulti={true}
											onChange={(val) => {
												onChange(
													speeds
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={speeds.find((c) => c.value === value)}
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

export default RAMFilter;

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
import * as hddAPI from "../../APIs/ProductMicroservice/hard_disk_drive_api";

const HDDFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [capacities, setCapacities] = useState([]);
	const [forms, setForms] = useState([]);
	const [diskSpeeds, setDiskSpeeds] = useState([]);

	useEffect(() => {
		getManufacturers();
		getCapacities();
		getForms();
		getDiskSpeeds();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Capacities: [],
			Forms: [],
			DiskSpeeds: [],
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
			.get(hddAPI.GET_MANUFACTURERS)
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
			.get(hddAPI.GET_CAPACITIES)
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

	const getForms = () => {
		axios
			.get(hddAPI.GET_FORMS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setForms(temp);
				setValue("Forms", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getDiskSpeeds = () => {
		axios
			.get(hddAPI.GET_DISK_SPEEDS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setDiskSpeeds(temp);
				setValue("DiskSpeeds", res.data);
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
							<Label>Filter by form</Label>
							<Controller
								name={"Forms"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={forms}
											isMulti={true}
											onChange={(val) => {
												onChange(
													forms
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={forms.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by disk speed</Label>
							<Controller
								name={"DiskSpeeds"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={diskSpeeds}
											isMulti={true}
											onChange={(val) => {
												onChange(
													diskSpeeds
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={diskSpeeds.find((c) => c.value === value)}
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

export default HDDFilter;

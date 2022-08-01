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
import * as processorAPI from "../../APIs/ProductMicroservice/processor_api";

const ProcessorFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [sockets, setSockets] = useState([]);
	const [types, setTypes] = useState([]);
	const [threads, setThreads] = useState([]);
	const [numberOfCores, setNumberOfCores] = useState([]);

	useEffect(() => {
		getManufacturers();
		getSockets();
		getTypes();
		getThreads();
		getNumberOfCores();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Sockets: [],
			Types: [],
			Threads: [],
			NumberOfCores: [],
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
			.get(processorAPI.GET_MANUFACTURERS)
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
			.get(processorAPI.GET_SOCKETS)
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

	const getTypes = () => {
		axios
			.get(processorAPI.GET_TYPES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setTypes(temp);
				setValue("Types", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getThreads = () => {
		axios
			.get(processorAPI.GET_THREADS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setThreads(temp);
				setValue("Threads", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getNumberOfCores = () => {
		axios
			.get(processorAPI.GET_NUMBER_OF_CORES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setNumberOfCores(temp);
				setValue("NumberOfCores", res.data);
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
							<Label>Filter by type</Label>
							<Controller
								name={"Types"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={types}
											isMulti={true}
											onChange={(val) => {
												onChange(
													types
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={types.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by number of threads</Label>
							<Controller
								name={"Threads"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={threads}
											isMulti={true}
											onChange={(val) => {
												onChange(
													threads
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={threads.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by number of cores</Label>
							<Controller
								name={"NumberOfCores"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={numberOfCores}
											isMulti={true}
											onChange={(val) => {
												onChange(
													numberOfCores
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={numberOfCores.find((c) => c.value === value)}
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

export default ProcessorFilter;

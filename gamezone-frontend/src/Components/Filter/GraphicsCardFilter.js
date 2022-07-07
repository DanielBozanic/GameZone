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
import * as graphicsCardsAPI from "../../APIs/ProductMicroservice/graphics_card_api";

const GraphicsCardFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [chipManufacturers, setChipManufacturers] = useState([]);
	const [memorySizes, setMemorySizes] = useState([]);
	const [memoryTypes, setMemoryTypes] = useState([]);
	const [modelNames, setModelNames] = useState([]);

	useEffect(() => {
		getManufacturers();
		getChipManufacturers();
		getMemorySizes();
		getMemoryTypes();
		getModelNames();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			ChipManufacturers: [],
			MemorySizes: [],
			MemoryTypes: [],
			ModelNames: [],
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
			.get(graphicsCardsAPI.GET_MANUFACTURERS)
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

	const getChipManufacturers = () => {
		axios
			.get(graphicsCardsAPI.GET_CHIP_MANUFACTURERS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setChipManufacturers(temp);
				setValue("ChipManufacturers", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getMemorySizes = () => {
		axios
			.get(graphicsCardsAPI.GET_MEMORY_SIZES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setMemorySizes(temp);
				setValue("MemorySizes", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getMemoryTypes = () => {
		axios
			.get(graphicsCardsAPI.GET_MEMORY_TYPES)
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

	const getModelNames = () => {
		axios
			.get(graphicsCardsAPI.GET_MODEL_NAMES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setModelNames(temp);
				setValue("ModelNames", res.data);
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
							<Label>Filter by chip manufacturer</Label>
							<Controller
								name={"ChipManufacturers"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={chipManufacturers}
											isMulti={true}
											onChange={(val) => {
												onChange(
													chipManufacturers
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={chipManufacturers.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by memory size</Label>
							<Controller
								name={"MemorySizes"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={memorySizes}
											isMulti={true}
											onChange={(val) => {
												onChange(
													memorySizes
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={memorySizes.find((c) => c.value === value)}
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
							<Label>Filter by model name</Label>
							<Controller
								name={"ModelNames"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={modelNames}
											isMulti={true}
											onChange={(val) => {
												onChange(
													modelNames
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={modelNames.find((c) => c.value === value)}
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

export default GraphicsCardFilter;

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
import * as monitorAPI from "../../APIs/ProductMicroservice/monitor_api";

const MonitorFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [aspectRatios, setAspectRatios] = useState([]);
	const [resolutions, setResolutions] = useState([]);
	const [refreshRates, setRefreshRates] = useState([]);

	useEffect(() => {
		getManufacturers();
		getAspectRatios();
		getResolutions();
		getRefreshRates();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			AspectRatios: [],
			Resolutions: [],
			RefreshRates: [],
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
			.get(monitorAPI.GET_MANUFACTURERS)
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

	const getAspectRatios = () => {
		axios
			.get(monitorAPI.GET_ASPECT_RATIOS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setAspectRatios(temp);
				setValue("AspectRatios", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getResolutions = () => {
		axios
			.get(monitorAPI.GET_RESOLUTIONS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setResolutions(temp);
				setValue("Resolutions", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getRefreshRates = () => {
		axios
			.get(monitorAPI.GET_REFRESH_RATES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setRefreshRates(temp);
				setValue("RefreshRates", res.data);
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
							<Label>Filter by aspect ratio</Label>
							<Controller
								name={"AspectRatios"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={aspectRatios}
											isMulti={true}
											onChange={(val) => {
												onChange(
													aspectRatios
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={aspectRatios.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by resolution</Label>
							<Controller
								name={"Resolutions"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={resolutions}
											isMulti={true}
											onChange={(val) => {
												onChange(
													resolutions
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={resolutions.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by refresh rate</Label>
							<Controller
								name={"RefreshRates"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={refreshRates}
											isMulti={true}
											onChange={(val) => {
												onChange(
													refreshRates
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={refreshRates.find((c) => c.value === value)}
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

export default MonitorFilter;

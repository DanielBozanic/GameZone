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
import * as videoGameAPI from "../../APIs/ProductMicroservice/video_game_api";

const VideoGamesFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [platforms, setPlatforms] = useState([]);
	const [genres, setGenres] = useState([]);

	useEffect(() => {
		getPlatforms();
		getGenres();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Platforms: [],
			Genres: [],
		},
	});

	const onFilterClick = (filterData) => {
		toggle();
		props.onFilterClick && props.onFilterClick(filterData);
	};

	const filter = (filterData) => {
		onFilterClick(filterData);
	};

	const getPlatforms = () => {
		axios
			.get(videoGameAPI.GET_PLATFORMS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setPlatforms(temp);
				setValue("Platforms", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getGenres = () => {
		axios
			.get(videoGameAPI.GET_GENRES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setGenres(temp);
				setValue("Genres", res.data);
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
							<Label>Filter by platform</Label>
							<Controller
								name={"Platforms"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={platforms}
											isMulti={true}
											onChange={(val) => {
												onChange(
													platforms
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={platforms.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>

						<FormGroup>
							<Label>Filter by genre</Label>
							<Controller
								name={"Genres"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={genres}
											isMulti={true}
											onChange={(val) => {
												onChange(
													genres
														.filter((o) => val.includes(o))
														.map((g) => g.value)
												);
											}}
											value={genres.find((c) => c.value === value)}
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

export default VideoGamesFilter;

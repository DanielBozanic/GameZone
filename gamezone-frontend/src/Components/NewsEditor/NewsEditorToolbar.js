import React from "react";
import { Quill } from "react-quill";

const CustomUndo = () => (
	<svg viewBox="0 0 18 18">
		<polygon className="ql-fill ql-stroke" points="6 10 4 12 2 10 6 10" />
		<path
			className="ql-stroke"
			d="M8.09,13.91A4.6,4.6,0,0,0,9,14,5,5,0,1,0,4,9"
		/>
	</svg>
);

const CustomRedo = () => (
	<svg viewBox="0 0 18 18">
		<polygon className="ql-fill ql-stroke" points="12 10 14 12 16 10 12 10" />
		<path
			className="ql-stroke"
			d="M9.91,13.91A4.6,4.6,0,0,1,9,14a5,5,0,1,1,5-5"
		/>
	</svg>
);

function undoChange() {
	this.quill.history.undo();
}

function redoChange() {
	this.quill.history.redo();
}

function imageHandler() {
	const range = this.quill.getSelection();
	const value = prompt("Please copy paste the image url here.");
	if (value) {
		this.quill.insertEmbed(range.index, "image", value, Quill.sources.USER);
	}
}

const Size = Quill.import("attributors/style/size");
const Align = Quill.import("attributors/style/align");
Size.whitelist = [
	"2px",
	"4px",
	"6px",
	"8px",
	"10px",
	"12px",
	"14px",
	"16px",
	"18px",
	"20px",
	"22px",
	"24px",
	"48px",
	"72px",
];
Quill.register(Size, true);
Quill.register(Align, true);

const Font = Quill.import("attributors/class/font");
Font.whitelist = [
	"arial",
	"comic-sans",
	"courier-new",
	"georgia",
	"helvetica",
	"monospace",
];
Quill.register(Font, true);

export const modules = {
	toolbar: {
		container: "#toolbar",
		handlers: {
			undo: undoChange,
			redo: redoChange,
			image: imageHandler,
		},
	},
	history: {
		delay: 500,
		maxStack: 100,
		userOnly: true,
	},
};

export const formats = [
	"header",
	"font",
	"size",
	"bold",
	"italic",
	"underline",
	"align",
	"strike",
	"script",
	"blockquote",
	"list",
	"bullet",
	"indent",
	"link",
	"image",
	"code-block",
];

export const NewsEditorToolbar = () => (
	<div id="toolbar">
		<span className="ql-formats">
			<select className="ql-font">
				<option value="arial">Arial</option>
				<option value="comic-sans">Comic Sans</option>
				<option value="courier-new">Courier New</option>
				<option value="georgia">Georgia</option>
				<option value="helvetica">Helvetica</option>
				<option value="monospace">Monospace</option>
			</select>
			<select className="ql-size">
				<option selected>Default</option>
				<option value="2px">2px</option>
				<option value="4px">4px</option>
				<option value="6px">6px</option>
				<option value="8px">8px</option>
				<option value="10px">10px</option>
				<option value="12px">12px</option>
				<option value="14px">14px</option>
				<option value="16px">16px</option>
				<option value="18px">18px</option>
				<option value="20px">20px</option>
				<option value="22px">22px</option>
				<option value="24px">24px</option>
				<option value="48px">48px</option>
				<option value="72px">72px</option>
			</select>
			<select className="ql-header">
				<option value="1">Heading</option>
				<option value="2">Subheading</option>
				<option value="3">Normal</option>
			</select>
		</span>
		<span className="ql-formats">
			<button className="ql-bold" />
			<button className="ql-italic" />
			<button className="ql-underline" />
			<button className="ql-strike" />
		</span>
		<span className="ql-formats">
			<button className="ql-list" value="ordered" />
			<button className="ql-list" value="bullet" />
			<button className="ql-indent" value="-1" />
			<button className="ql-indent" value="+1" />
		</span>
		<span className="ql-formats">
			<button className="ql-script" value="super" />
			<button className="ql-script" value="sub" />
			<button className="ql-blockquote" />
			<button className="ql-direction" />
		</span>
		<span className="ql-formats">
			<select className="ql-align" />
		</span>
		<span className="ql-formats">
			<button className="ql-link" />
			<button className="ql-image" />
		</span>
		<span className="ql-formats">
			<button className="ql-code-block" />
			<button className="ql-clean" />
		</span>
		<span className="ql-formats">
			<button className="ql-undo">
				<CustomUndo />
			</button>
			<button className="ql-redo">
				<CustomRedo />
			</button>
		</span>
	</div>
);

export default NewsEditorToolbar;

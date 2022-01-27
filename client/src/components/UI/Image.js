import "./Image.css";

const Image = (props) => {
  // Get the hard-coded height passed through props
  const adjustedHeight = props.HEIGHT;

  // Calculate the adjusted width to keep image proportional
  // let extraHeight = 0;
  // if (props.height > adjustedHeight)
  //   extraHeight = props.height - adjustedHeight;
  // else extraHeight = adjustedHeight - props.height;
  // const adjustRatio = extraHeight / props.height;
  // const adjustedWidth = (1 - adjustRatio) * props.width;

  return (
    <div>
      <img src={props.src} height={adjustedHeight} />
    </div>
  );
};

export default Image;

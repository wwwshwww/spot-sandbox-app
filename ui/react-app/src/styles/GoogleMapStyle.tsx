export { mapStyles, mapOptions };

const mapStyles = {
  // margin: 20,
  width: document.documentElement.clientWidth * 0.5,
  height: document.documentElement.clientWidth * 0.5,
  minWidth: 400,
  minHeight: 400,
};

const mapOptions = {
  disableDoubleClickZoom: true,
  draggableCursor: 'crosshair',
  clickableIcons: false,
  styles: [
    {
      stylers: [
        {
          saturation: -10
        },
        {
          lightness: 60
        }
      ]
    },
    {
      featureType: 'poi',
      elementType: 'labels.icon',
      stylers: [
        {
          lightness: 60
        }
      ]
    },
  ],
};

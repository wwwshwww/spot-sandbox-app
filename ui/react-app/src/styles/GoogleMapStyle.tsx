export { mapStyles, mapOptions };

const mapStyles = {
  // margin: 20,
  width: document.documentElement.clientWidth * 0.7,
  height: document.documentElement.clientWidth * 0.7,
  minWidth: 500,
  minHeight: 500,
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

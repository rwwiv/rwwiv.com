const importImage = (folder, filename) => {
  return require(`../assets/img/${folder}/${filename}`)
}

export  { importImage }
import request from '@/utils/request'

export function createProofMaterial(data) {
  return request({
    url: '/createProofMaterial',
    method: 'post',
    data
  })
}

export function queryProofCertify(data) {
  return request({
    url: '/queryProofCertify',
    method: 'post',
    data
  })
}

export function queryProofMaterial(data) {
  return request({
    url: '/queryProofMaterial',
    method: 'post',
    data
  })
}

export function createPhotoMaterial(data) {
  return request({
    url: '/createPhotoMaterial',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'multipart/form-data;'
    }
  })
}

export function queryPhotoMaterial(data) {
  return request({
    url: '/queryPhotoMaterial',
    method: 'post',
    data
  })
}

export function queryProofMaterialOnly(data) {
  return request({
    url: '/queryProofMaterialOnly',
    method: 'post',
    data
  })
}

export function updateProofMaterial(data) {
  return request({
    url: '/updateProofMaterial',
    method: 'post',
    data
  })
}


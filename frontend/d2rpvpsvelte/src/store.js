// store.js
import { writable } from 'svelte/store';

export const mode = writable(false);
export const nickname = writable(null);
export const formopen = writable(null);

export const csrfToken = writable(null);

export async function getCsrfToken() {
    const response = await fetch('/csrf-token');
    const data = await response.json();
    csrfToken.set(data.csrfToken);
    return data.csrfToken;
}

export function clickformopen(whatform) {
  formopen.set(whatform);
}


getCsrfToken();


// 세션 유효성 및 사용자 정보를 저장할 스토어
export const isAuthenticated = writable(false);

export const jwtoken = writable('');


export async function checkSession() {
    try {
      const response = await fetch("/check-session", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },      credentials: 'include'  // 쿠키 포함

      });
  
      if (response.ok) {
        const data = await response.json();
        isAuthenticated.set(data.authenticated);
        nickname.set(data.user.nickname);
        return data.authenticated;
      } else {
        isAuthenticated.set(false);
        nickname.set(null);
        return false;
      }
    } catch (error) {
      console.error("Error checking session:", error);
      isAuthenticated.set(false);
      nickname.set(null);
      return false;
    }
  }
  
  
 

// 닉네임 리스트를 저장할 store
export const nicknames = writable([]);

// 닉네임을 서버에서 가져오는 함수
export async function fetchNicknames(modevalue) {
  try {
    const endpoint = modevalue ? "/api/getNicknames_m" : "/api/getNicknames";
    const response = await fetch(endpoint);
    if (response.ok) {
      const data = await response.json();
      nicknames.set([]);
      nicknames.update(current => [...current, ...data]);
    } else {
      console.error('닉네임을 가져오는데 실패했습니다.');
    }
  } catch (error) {
    console.error('닉네임을 가져오는 중 오류가 발생했습니다:', error);
  }
}


export const key = writable(0);


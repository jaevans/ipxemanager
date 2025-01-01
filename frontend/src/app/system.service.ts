import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SystemService {
  private apiUrl = '/api';

  constructor(private http: HttpClient) {}

  getSystems(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/systems`);
  }

  getConfigurations(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/configurations`);
  }

  addSystem(system: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/systems`, system);
  }

  addConfiguration(configuration: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/configurations`, configuration);
  }

  updateSystem(system: any): Observable<any> {
    return this.http.put<any>(`${this.apiUrl}/systems/${system.id}`, system);
  }

  updateConfiguration(configuration: any): Observable<any> {
    return this.http.put<any>(`${this.apiUrl}/configurations/${configuration.id}`, configuration);
  }

  deleteSystem(systemId: number): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/systems/${systemId}`);
  }

  deleteConfiguration(configurationId: number): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/configurations/${configurationId}`);
  }

  getOneTimeOverrides(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/one-time-overrides`);
  }

  addOneTimeOverride(override: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/one-time-overrides`, override);
  }

  deleteOneTimeOverride(overrideId: number): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/one-time-overrides/${overrideId}`);
  }
}

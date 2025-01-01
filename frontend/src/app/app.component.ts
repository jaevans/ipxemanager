import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  systems: any[] = [];
  configurations: any[] = [];
  newSystem: any = {};
  newConfiguration: any = {};

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.fetchSystems();
    this.fetchConfigurations();
  }

  fetchSystems() {
    this.http.get('/api/systems').subscribe((data: any) => {
      this.systems = data;
    });
  }

  fetchConfigurations() {
    this.http.get('/api/configurations').subscribe((data: any) => {
      this.configurations = data;
    });
  }

  addSystem() {
    this.http.post('/api/systems', this.newSystem).subscribe(() => {
      this.fetchSystems();
      this.newSystem = {};
    });
  }

  addConfiguration() {
    this.http.post('/api/configurations', this.newConfiguration).subscribe(() => {
      this.fetchConfigurations();
      this.newConfiguration = {};
    });
  }

  updateSystem(system: any) {
    this.http.put(`/api/systems/${system.id}`, system).subscribe(() => {
      this.fetchSystems();
    });
  }

  updateConfiguration(configuration: any) {
    this.http.put(`/api/configurations/${configuration.id}`, configuration).subscribe(() => {
      this.fetchConfigurations();
    });
  }

  deleteSystem(system: any) {
    this.http.delete(`/api/systems/${system.id}`).subscribe(() => {
      this.fetchSystems();
    });
  }

  deleteConfiguration(configuration: any) {
    this.http.delete(`/api/configurations/${configuration.id}`).subscribe(() => {
      this.fetchConfigurations();
    });
  }
}
